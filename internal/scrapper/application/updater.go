package application

import (
	"context"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
	"log/slog"
	"time"
)

type AllLinksGetter interface {
	GetAllLinks(ctx context.Context) ([]domain.TrackedLink, error)
}

type Updater struct {
	links    AllLinksGetter
	checkers []LinkChecker
	logger   *slog.Logger
	period   time.Duration
}

func NewUpdater(links AllLinksGetter, checkers []LinkChecker, logger *slog.Logger, period time.Duration) *Updater {
	return &Updater{
		links:    links,
		checkers: checkers,
		logger:   logger,
		period:   period,
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
		var checker LinkChecker

		for _, ch := range u.checkers {
			if ch.Supports(link.URL) {
				checker = ch
				break
			}
		}

		if checker == nil {
			u.logger.Warn("no checker for link", "url", link.URL)
			continue
		}

		result, err := checker.Check(ctx, link)
		if err != nil {
			u.logger.Error("failed to check link", "url", link.URL, "error", err)
			continue
		}

		if result.HasUpdates {
			u.logger.Info("link updated",
				"url", link.URL,
				"description", result.Description,
				"updated_at", result.NewUpdatedAt,
			)
		} else {
			u.logger.Info("no updates", "url", link.URL)
		}
	}
}
