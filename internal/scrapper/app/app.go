package app

import (
	"log/slog"
	nethttp "net/http"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/config"
	httpserver "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/infrastructure/http"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/infrastructure/memory"
)

type App struct {
	config *config.Config
	logger *slog.Logger
	server *nethttp.Server
}


func New(cfg *config.Config, logger *slog.Logger) *App{
	chats := memory.NewChatRepository()
	service := application.NewService(chats)
	server := httpserver.NewServer(service, cfg)

	return &App{
		config: cfg,
		logger: logger,
		server: server,
	}
}

func (a *App) Run() error {
	a.logger.Info("starting scrapper", "port", a.config.Port)
	return a.server.ListenAndServe()
}