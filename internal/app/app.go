package app

import (
	"context"
	"log/slog"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/dispatcher"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/telegram"
)

type App struct {
	token string
	log *slog.Logger
	telegram *telegram.TelegramBot
	dispatcher *dispatcher.Dispatcher
}


func New(cfg *config.Config, logger *slog.Logger, telegram *telegram.TelegramBot) *App {
	d := dispatcher.New()
	telegram.RegistrationCommands(d);

	return &App {
		token: cfg.TelegramToken,
		log: logger,
		telegram: telegram,
		dispatcher: d,
	}
}


func (a *App) Run(ctx context.Context) error {
	bot := a.telegram;
	api := a.telegram.Api;

	if err := bot.SetMyCommands(a.dispatcher); err != nil {
		a.log.Info("mycommands not register in tg_bot",
		"error", err)
	}

	events := bot.ReceiveMessages()

	for {
		select {
		case <- ctx.Done():
			api.StopReceivingUpdates()
			return nil;
		case event, ok := <- events:
			if !ok {
				return nil
			}

			if event.Type != "command" {
				continue
			}

			text := event.Message.Text;
			cmd := event.Message.Command;
			chatID := event.Message.ChatID;

			a.log.Info("update received",
				"chat_id", chatID,
				"text", text,
			)

			reply, ok := dispatcher.Dispatch(a.dispatcher, cmd);
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
				return err
			}
		}
	}
}