package app

import (
	"context"
	"log/slog"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/dispatcher"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/telegram"
)

type App struct {
	token      string
	log        *slog.Logger
	telegram   *telegram.TelegramBot
	dispatcher *dispatcher.Dispatcher
}

func New(cfg *config.Config, logger *slog.Logger, telegram *telegram.TelegramBot) *App {
	d := dispatcher.New()

	descriptions := map[command.Name]string{
		command.Start: "запуск телеграмм бота",
		command.Help:  "список доступных команд",
	}

	startcmd := handler.Start{}
	helpcmd := handler.Help{
		Descriptions: descriptions,
	}

	d.Register(command.Start, descriptions[command.Start], startcmd)
	d.Register(command.Help, descriptions[command.Help], helpcmd)

	return &App{
		token:      cfg.TelegramToken,
		log:        logger,
		telegram:   telegram,
		dispatcher: d,
	}
}

func (a *App) Run(ctx context.Context) error {
	bot := a.telegram

	if err := bot.SetCommands(a.dispatcher); err != nil {
		a.log.Info("mycommands not register in tg_bot",
			"error", err)
	}

	events := bot.ReceiveMessages(ctx)

	for event := range events {
		text := event.Text
		cmd := event.Command
		chatID := event.ChatID

		a.log.Info("update received",
			"chat_id", chatID,
			"text", text,
		)

		reply, ok := a.dispatcher.Dispatch(command.Name(cmd))
		if !ok {
			continue
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
