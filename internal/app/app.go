package app

import (
	"context"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/dispatcher"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handlers"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/outer"
)

type App struct {
	token string
	log *slog.Logger
	dispatcher *dispatcher.Dispatcher
}


func New(cfg config.Config, logger *slog.Logger) *App {
	d := dispatcher.New()
	d.Register("start", "запуск телеграмм бота", handlers.Start)
	getHelpText := func() string {return outer.HelpText(d)}
	d.Register("help", "список доступных команд", handlers.Help(func() string {return getHelpText()}))

	return &App {
		log: logger,
		token: cfg.TelegramToken,
		dispatcher: d,
	}
}


func (a *App) Run(ctx context.Context) error {
	bot, err := tgbotapi.NewBotAPI(a.token);
	if err != nil {
		a.log.Error("failed on starting bot", 
		"error", err);
		return err;
	}

	if err := outer.SetMyCommands(bot, a.dispatcher); err != nil {
		a.log.Info("mycommands not register in tg_bot",
		"error", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30; //удержание запроса

	updates := bot.GetUpdatesChan(u);

	for {
		select {
		case <- ctx.Done():
			bot.StopReceivingUpdates()
			return nil;
		case update, ok := <- updates:
			if !ok {
				return nil
			}

			if update.Message == nil {
				continue;
			}

			if !update.Message.IsCommand() { 
				continue 
			}

			text := update.Message.Text;
			cmd := update.Message.Command();
			chat_id := update.Message.Chat.ID;

			a.log.Info("update received",
				"chat_id", chat_id,
				"text", text,
			)

			reply, ok := outer.Dispatch(a.dispatcher, cmd);
			if !ok {
				continue	
			}

			a.log.Info("reply prepared",
				"chat_id", chat_id,
				"reply", reply,
			)

			resp := tgbotapi.NewMessage(chat_id, reply);
			if _, err := bot.Send(resp); err != nil {
				a.log.Error("send failed",
					"chat_id", chat_id,
					"err", err,
				)
				return err;
			}
		}
	}
}