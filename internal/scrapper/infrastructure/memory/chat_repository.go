package memory

import (
	"context"
	"fmt"
	"sync"
)

type ChatRepository struct {
	mu    sync.RWMutex
	chats map[int64]struct{}
}


func NewChatRepository() *ChatRepository {
	return &ChatRepository{
		chats: make(map[int64]struct{}),
	}
}


func (r *ChatRepository) Add(ctx context.Context, chatID int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.chats[chatID]; exists {
		return fmt.Errorf("chat already exists")
	}

	r.chats[chatID] = struct{}{}
	return nil
}


func (r *ChatRepository) Remove(ctx context.Context, chatID int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.chats[chatID]; !exists {
		return fmt.Errorf("chat not found")
	}

	delete(r.chats, chatID)
	return nil
}

