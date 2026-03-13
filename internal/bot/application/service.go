package application

import "context"

type Service struct {
	scrapper ScrapperClient
}

func NewService(scrapper ScrapperClient) *Service {
	return &Service{
		scrapper: scrapper,
	}
}

func (s *Service) RegisterChat(ctx context.Context, chatID int64) error {
	return s.scrapper.RegisterChat(ctx, chatID)
}