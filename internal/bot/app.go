package app

import (
	"context"
	"log/slog"
	"net/http"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/dispatcher"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/telegram"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/infrastructure/client"
)

type App struct {
	token      string
	log        *slog.Logger
	telegram   *telegram.TelegramBot
	dispatcher *dispatcher.Dispatcher
	botService *application.Service
}


func New(cfg *config.Config, logger *slog.Logger, telegram *telegram.TelegramBot) *App {
	d := dispatcher.New()

	httpClient := &http.Client{}

	scrapperClient := client.New(cfg.BaseUrl+cfg.Port, httpClient)

	botService := application.NewService(scrapperClient)
	
	return &App{
		token:      cfg.TelegramToken,
		log:        logger,
		telegram:   telegram,
		dispatcher: d,
		botService: botService,
	}
}

func (a *App) Run(ctx context.Context) error {
	bot := a.telegram

	descriptions := map[domain.Name]string{
		domain.Start: "запуск телеграмм бота",
		domain.Help:  "список доступных команд",
		domain.Track: "добавить отслеживание ссылки",
		domain.Cancel: "отмена добавления ссылки",
		domain.List: "Список отслеживаемых ссылок",
		domain.Untrack: "Удаление ссылки",
	}

	startcmd := handler.Start{
		Logger: a.log,
		Service: a.botService,
	}

	helpcmd := handler.Help{
		Descriptions: descriptions,
	}

	trackcmd := handler.Track{
		Service: a.botService,
	}

	cancelcmd := handler.Cancel{
		Service: a.botService,
	}

	listcmd := handler.List{
		Service: a.botService,
	}

	untrackcmd := handler.Untrack {
		Service: a.botService,
	}

	a.dispatcher.Register(domain.Start, descriptions[domain.Start], startcmd)
	a.dispatcher.Register(domain.Help, descriptions[domain.Help], helpcmd)
	a.dispatcher.Register(domain.Track, descriptions[domain.Track], trackcmd)
	a.dispatcher.Register(domain.Cancel, descriptions[domain.Cancel], cancelcmd)
	a.dispatcher.Register(domain.List, descriptions[domain.List], listcmd)
	a.dispatcher.Register(domain.Untrack, descriptions[domain.Untrack], untrackcmd)

	if err := bot.SetCommands(descriptions); err != nil {
		a.log.Info("mycommands not register in tg_bot", "error", err)
	}

	events := bot.ReceiveMessages(ctx)

	for event := range events {
		text := event.Text
		cmd := event.Command
		chatID := event.ChatID

		req := domain.Request{
			Context: ctx,
			ChatID: chatID,
			Text: text,
		}

		a.log.Info("update received",
			"chat_id", chatID,
			"text", text,
		)

		var reply string

		if a.botService.HasActiveTrackDialog(chatID) {
			if cmd == "" {
				reply = a.botService.ContinueTrackDialog(ctx, chatID, text)
			} else if cmd == string(domain.Cancel) {
				reply = a.dispatcher.Dispatch(domain.Name(cmd), req)
			} else {
				reply = "Сначала завершите заполнение команды /track или отправьте /cancel"
			}
		} else {
			reply = a.dispatcher.Dispatch(domain.Name(cmd), req)
		}

		a.log.Info("reply prepared",
			"chat_id", chatID,
			"reply", reply,
		)

		if err := bot.Send(chatID, reply); err != nil {
			a.log.Error("send failed",
				"chat_id", chatID,
				"err", err,
			)
		}
	}

	return nil
}
