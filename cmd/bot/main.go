package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/logger"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/telegram"
)

func main() {
	cfg, err := config.Load("./cmd/bot/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	logger := logger.New(cfg.LogLevel)

	telegram, err := telegram.NewTelegramBot(cfg.TelegramToken)

	if err != nil {
		logger.Error("bot not created",
			"error", err)
	}
	App := app.New(cfg, logger, telegram)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt) //перестаем крутить Run() после ctrl+c
	defer stop()

	if err := App.Run(ctx); err != nil {
		logger.Error("app stopped with error",
			"error", err)
	}
}
