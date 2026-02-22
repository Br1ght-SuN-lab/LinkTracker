package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/application/bot/dispatcher"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/application/bot/handlers"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/infrastructure/telegram"
)


func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	cfg, err := config.Load();
	if err != nil {
		log.Fatalf("failed on get token: %v", err);
	}

	fmt.Println("Bot starting with length of token:", len(cfg.TelegramToken));
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken);
	if err != nil {
		log.Fatalf("failed on creating bot:%v", err);
	}

	if err := telegram.RegisterCommands(bot); err != nil {
		log.Printf("failed to register commands: %v", err)
	}
	u := tgbotapi.NewUpdate(0);
	u.Timeout = 30; //удержание запроса

	updates := bot.GetUpdatesChan(u);
	for update := range updates {
		if update.Message == nil {
			continue;
		}

		text := update.Message.Text;
		chat_id := update.Message.Chat.ID;

		logger.Info("update received",
			"chat_id", chat_id,
			"text", text,
		)

		d := dispatcher.New();

		//в одну строчку можно добавить новую команду
		d.Register("/start", handlers.StartHandler)
		d.Register("/help", handlers.HelpHandler)

		reply, ok := d.Dispatch(text);
		if !ok {
			continue
		}

		logger.Info("reply prepared",
			"chat_id", chat_id,
			"reply", reply,
		)

		resp := tgbotapi.NewMessage(chat_id, reply);
		if _, err := bot.Send(resp); err != nil {
			logger.Error("send failed",
				"chat_id", chat_id,
				"err", err,
			)
		}
	}
}