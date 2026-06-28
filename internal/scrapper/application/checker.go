package application

import (
	"context"
	"time"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)

type CheckResult struct {
	HasUpdates   bool
	Description  string
	NewUpdatedAt time.Time
}

type LinkChecker interface {
	Supports(rawURL string) bool
	Check(ctx context.Context, link domain.TrackedLink) (CheckResult, error)
}
