package main

import (
	"log"
	"log/slog"
	"os"

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

	if err := App.Run(); err != nil {
		logger.Error("scrapper not listen", "error", err)
	}
	
}