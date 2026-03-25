package handler

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application"
)
type Track struct {
	Service *application.Service
}

func (h Track) Handle(req domain.Request) string {
	h.Service.StartTrackDialog(req.ChatID)
	return "Пришли ссылку, которую нужно отслеживать"
}

// func (h Track) Handle(req domain.Request) string {
// 	tokens := strings.Fields(req.Text)
// 	if len(tokens) < 2 {
// 		return "Используйте формат ввода: /track <url>"
// 	}

// 	url := tokens[1];

// 	if _, err := h.Service.AddLink(req.Context, req.ChatID, url, nil, nil); err != nil {
// 		return "Не удалось добавить ссылку: " + err.Error()
// 	}

// 	return "ссылка добавлена"
// }