package application

import "context"

type ChatRepository interface {
	Add(ctx context.Context, chatID int64) error 
	Remove(ctx context.Context, chatID int64) error 
}

type Service struct {
	chats ChatRepository
}


func NewService(chats ChatRepository) *Service {
	return &Service{
		chats: chats,
	}
}


func (s *Service) RegisterChat(ctx context.Context, chatID int64) error {
	return s.chats.Add(ctx, chatID)
}


func (s *Service) DeleteChat(ctx context.Context, chatID int64) error {
	return s.chats.Remove(ctx, chatID)
}