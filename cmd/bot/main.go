package main

import (
	"log"
	"context"
	"os"
	"os/signal"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/app"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/logger"
)


func main() {
	cfg, err := config.Load("config.yaml");
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err) //не придумал как без нагромождений логгировать
	}

	logger := logger.New(cfg.LogLevel)

	App := app.New(cfg, logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt) //перестаем крутить Run() после ctrl+c
	defer stop()

	if err := App.Run(ctx); err != nil {
		logger.Error("app stopped with error", 
		"error", err)
	}
}