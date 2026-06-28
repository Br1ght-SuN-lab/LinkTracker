package application

import (
	"context"
	"net/url"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/scrapper/domain"
)

type ScrapperClient interface {
	RegisterChat(ctx context.Context, chatID int64) error
	DeleteChat(ctx context.Context, chatID int64) error

	AddLink(ctx context.Context, chatID int64, url string, tags []string, filters []string) (domain.Link, error)
	GetLinks(ctx context.Context, chatID int64) ([]domain.Link, error)
	RemoveLink(ctx context.Context, chatID int64, url string) (domain.Link, error)
}

type Service struct {
	client ScrapperClient
	trackState *TrackStateStorage
}

func NewService(client ScrapperClient) *Service {
	return &Service{
		client: client,
		trackState: NewTrackStateStorage(),
	}
}

//Client 
func (s *Service) RegisterChat(ctx context.Context, chatID int64) error {
	return s.client.RegisterChat(ctx, chatID)
}

func (s *Service) DeleteChat(ctx context.Context, chatID int64) error {
	return s.client.DeleteChat(ctx, chatID)
}

func (s *Service) AddLink(ctx context.Context, chatID int64, url string, tags []string, filters []string) (domain.Link, error) {
	return s.client.AddLink(ctx, chatID, url, tags, filters)
}

func (s *Service) ListLinks(ctx context.Context, chatID int64) ([]domain.Link, error) {
	return s.client.GetLinks(ctx, chatID)
}

func (s *Service) RemoveLink(ctx context.Context, chatID int64, url string) (domain.Link, error) {
	return s.client.RemoveLink(ctx, chatID, url)
}

//TrackDialog
func (s *Service) StartTrackDialog(chatID int64) {
	s.trackState.Start(chatID)
}

func (s *Service) GetTrackDialog(chatID int64) (TrackDialog, bool) {
	return s.trackState.Get(chatID)
}

func (s *Service) UpdateTrackDialog(chatID int64, state TrackDialog) {
	s.trackState.Update(chatID, state)
}

func (s *Service) DeleteTrackDialog(chatID int64) {
	s.trackState.Delete(chatID)
}

func (s *Service) HasActiveTrackDialog(chatID int64) bool {
	return s.trackState.Exists(chatID)
}


func (s *Service) ContinueTrackDialog(ctx context.Context, chatID int64, text string) string {
	state, ok := s.trackState.Get(chatID)
	if !ok {
		return ""
	}

	switch state.Step {
	case TrackStepLink:
		if !isValidURL(text) {
			return "Некорректная ссылка. Пришли корректный URL"
		}

		state.Link = text
		state.Step = TrackStepTags
		s.trackState.Update(chatID, state)
		return "Теперь пришли теги через пробел. Если тегов нет — отправь -"

	case TrackStepTags:
		if text == "-" {
			state.Tags = nil
		} else {
			state.Tags = strings.Fields(text)
		}

		state.Step = TrackStepFilters
		s.trackState.Update(chatID, state)
		return "Теперь пришли фильтры через пробел. Если фильтров нет — отправь -"

	case TrackStepFilters:
		if text == "-" {
			state.Filters = nil
		} else {
			state.Filters = strings.Fields(text)
		}

		_, err := s.client.AddLink(ctx, chatID, state.Link, state.Tags, state.Filters)
		s.trackState.Delete(chatID)

		if err != nil {
			return "Не удалось добавить ссылку: " + err.Error()
		}

		return "Ссылка успешно добавлена"

	default:
		s.trackState.Delete(chatID)
		return "Диалог сброшен. Попробуй заново: /track"
	}
}

func isValidURL(raw string) bool {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}

	return u.Scheme != "" && u.Host != ""
}