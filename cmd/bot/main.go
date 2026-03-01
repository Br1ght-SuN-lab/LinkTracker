package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/app"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
)


func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	cfg, err := config.Load();
	if err != nil {
		logger.Error("failed on get token", 
		"error", err);
	}

	App := app.New(cfg, logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt) //перестаем крутить Run() после ctrl+c
	defer stop()

	if err := App.Run(ctx); err != nil {
		logger.Error("app stopped with error", 
		"error", err)
	}
}