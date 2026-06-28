package app

import (
	"context"
	"errors"
	"log/slog"
	nethttp "net/http"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/config"
	githubchecker "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/infrastructure/checker/github"
	httpserver "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/infrastructure/http"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/infrastructure/memory"
)

type App struct {
	config  *config.Config
	logger  *slog.Logger
	server  *nethttp.Server
	updater *application.Updater
}

func New(cfg *config.Config, logger *slog.Logger) *App {
	chats := memory.NewChatRepository()
	links := memory.NewLinkRepository()
	service := application.NewService(chats, links)
	server := httpserver.NewServer(service, cfg)

	checkers := []application.LinkChecker{
		githubchecker.New(nil, cfg.GithubToken),
	}
	updater := application.NewUpdater(links, checkers, logger, 30*time.Second)

	return &App{
		config:  cfg,
		logger:  logger,
		server:  server,
		updater: updater,
	}
}

func (a *App) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go a.updater.Run(ctx)

	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, nethttp.ErrServerClosed) {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		a.logger.Info("shutting down application")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return a.server.Shutdown(shutdownCtx)

	case err := <-errCh:
		return err
	}
}
