package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	app "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/app"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/config"
)

func main() {
	cfg, err := config.Load("./cmd/scrapper/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	App := app.New(cfg, logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := App.Run(ctx); err != nil {
		logger.Error("application stopped with error", "error", err)
		os.Exit(1)
	}

}
