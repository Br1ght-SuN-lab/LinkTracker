package application

import (
	"context"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)

type ChatRepository interface {
	Add(ctx context.Context, chatID int64) error
	Remove(ctx context.Context, chatID int64) error
}

type LinksRepository interface {
	AddLink(ctx context.Context, chatID int64, url string, tags []string, filters []string) (domain.Link, error)
	GetLinks(ctx context.Context, chatID int64) ([]domain.Link, error)
	RemoveLink(ctx context.Context, chatID int64, url string) (domain.Link, error)
	GetAllLinks(ctx context.Context) ([]domain.TrackedLink, error)
}


type Service struct {
	chats ChatRepository
	links LinksRepository
}

func NewService(chats ChatRepository, links LinksRepository) *Service {
	return &Service{
		chats: chats,
		links: links,
	}
}

func (s *Service) RegisterChat(ctx context.Context, chatID int64) error {
	return s.chats.Add(ctx, chatID)
}

func (s *Service) DeleteChat(ctx context.Context, chatID int64) error {
	return s.chats.Remove(ctx, chatID)
}

func (s *Service) AddLink(ctx context.Context, chatID int64, url string, tags []string, filters []string) (domain.Link, error) {
	return s.links.AddLink(ctx, chatID, url, tags, filters)
}

func (s *Service) ListLinks(ctx context.Context, chatID int64) ([]domain.Link, error) {
	return s.links.GetLinks(ctx, chatID)
}

func (s *Service) RemoveLink(ctx context.Context, chatID int64, url string) (domain.Link, error) {
	return s.links.RemoveLink(ctx, chatID, url)
}

func (s *Service) GetAllLinks(ctx context.Context) ([]domain.TrackedLink, error) {
	return s.links.GetAllLinks(ctx)
}