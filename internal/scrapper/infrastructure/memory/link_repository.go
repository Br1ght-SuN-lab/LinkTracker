package memory

import (
	"context"
	"fmt"
	"sync"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)

type LinkRepository struct {
	mu     sync.RWMutex
	data   map[int64]map[string]domain.Link
	nextID int64
}

func NewLinkRepository() *LinkRepository {
	return &LinkRepository{
		data:   make(map[int64]map[string]domain.Link),
		nextID: 1,
	}
}

func (r *LinkRepository) AddLink(ctx context.Context, chatID int64, url string, tags []string, filters []string) (domain.Link, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[chatID]; !ok {
		r.data[chatID] = make(map[string]domain.Link)
	}

	if _, exists := r.data[chatID][url]; exists {
		return domain.Link{}, fmt.Errorf("link already exists")
	}

	link := domain.Link{
		ID:      r.nextID,
		URL:     url,
		Tags:    append([]string(nil), tags...),
		Filters: append([]string(nil), filters...),
	}

	r.data[chatID][url] = link
	r.nextID++

	return link, nil
}

func (r *LinkRepository) GetLinks(ctx context.Context, chatID int64) ([]domain.Link, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	linksMap, ok := r.data[chatID]
	if !ok {
		return []domain.Link{}, nil
	}

	links := make([]domain.Link, 0, len(linksMap))
	for _, link := range linksMap {
		links = append(links, link)
	}

	return links, nil
}

func (r *LinkRepository) RemoveLink(ctx context.Context, chatID int64, url string) (domain.Link, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	linksMap, ok := r.data[chatID]
	if !ok {
		return domain.Link{}, fmt.Errorf("link not found")
	}

	link, exists := linksMap[url]
	if !exists {
		return domain.Link{}, fmt.Errorf("link not found")
	}

	delete(linksMap, url)

	return link, nil
}


func (r *LinkRepository) GetAllLinks(ctx context.Context) ([]domain.TrackedLink, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.TrackedLink, 0)

	for chatID, chatLinks := range r.data {
		for _, link := range chatLinks {
			result = append(result, domain.TrackedLink{
				ChatID:  chatID,
				ID:      link.ID,
				URL:     link.URL,
				Tags:    append([]string(nil), link.Tags...),
				Filters: append([]string(nil), link.Filters...),
			})
		}
	}

	return result, nil
}