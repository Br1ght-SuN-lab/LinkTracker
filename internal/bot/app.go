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
	commands   map[command.Name]string
}

func New(cfg *config.Config, logger *slog.Logger, telegram *telegram.TelegramBot) *App {
	d := dispatcher.New()

	descriptions := map[command.Name]string{
		command.Start: "запуск телеграмм бота",
		command.Help:  "список доступных команд",
	}

	return &App{
		token:      cfg.TelegramToken,
		log:        logger,
		telegram:   telegram,
		dispatcher: d,
		commands: descriptions,
	}
}

func (a *App) Run(ctx context.Context) error {
	startcmd := handler.Start{}
	helpcmd := handler.Help{
		Descriptions: a.commands,
	}

	a.dispatcher.Register(startcmd)
	a.dispatcher.Register(helpcmd)

	bot := a.telegram

	if err := bot.SetCommands(a.commands); err != nil {
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

		reply:= a.dispatcher.Dispatch(command.Name(cmd))

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
