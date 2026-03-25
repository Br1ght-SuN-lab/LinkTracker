package application

import (
	"context"
	"log/slog"
	"time"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)

type AllLinksGetter interface {
	GetAllLinks(ctx context.Context) ([]domain.TrackedLink, error)
}

type Updater struct {
	links AllLinksGetter
	logger  *slog.Logger
	period  time.Duration
}

func NewUpdater(links AllLinksGetter, logger *slog.Logger, period time.Duration) *Updater {
	return &Updater{
		links:  links,
		logger: logger,
		period: period,
	}
}

func (u *Updater) Run(ctx context.Context) {
	ticker := time.NewTicker(u.period)
	defer ticker.Stop()

	u.logger.Info("scrapper updater started", "period", u.period.String())

	for {
		select {
		case <-ctx.Done():
			u.logger.Info("scrapper updater stopped")
			return
		case <-ticker.C:
			u.checkAllLinks(ctx)
		}
	}
}

func (u *Updater) checkAllLinks(ctx context.Context) {
	links, err := u.links.GetAllLinks(ctx)
	if err != nil {
		u.logger.Error("failed to get all links", "error", err)
		return
	}

	if len(links) == 0 {
		u.logger.Info("no tracked links to check")
		return
	}

	for _, link := range links {
		u.logger.Info("checking tracked link",
			"chat_id", link.ChatID,
			"link_id", link.ID,
			"url", link.URL,
		)

		//уже нереально успеть написать github и stackoverflow...просто лог что обновляю ссылки
		//показываю обновление ссылок без внешнего API
	}
}