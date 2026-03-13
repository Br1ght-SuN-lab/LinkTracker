package application

import (
	"context"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)

type ScrapperClient interface {
	RegisterChat(ctx context.Context, chatID int64) error
	DeleteChat(ctx context.Context, chatID int64) error
	GetLinks(ctx context.Context, chatID int64) ([]domain.Link, error)
	AddLink(ctx context.Context, chatID int64, link string, tags, filters []string) (domain.Link, error)
	RemoveLink(ctx context.Context, chatID int64, link string) (domain.Link, error)
}