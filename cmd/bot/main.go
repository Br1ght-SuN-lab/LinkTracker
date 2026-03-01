package main

import (
	"fmt"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/app"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/infrastructure/config"
)


func main() {
	cfg, err := config.Load();
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err) //не придумал как без нагромождений логгировать
		os.Exit(1)
	}

	level := config.ParseLogLevel(cfg.LogLevel)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	App := app.New(cfg, logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt) //перестаем крутить Run() после ctrl+c
	defer stop()

	if err := App.Run(ctx); err != nil {
		logger.Error("app stopped with error", 
		"error", err)
	}
}