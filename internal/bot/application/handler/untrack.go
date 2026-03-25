package handler

import (
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain"
)

type Untrack struct {
	Service *application.Service
}

func (h Untrack) Handle(req domain.Request) string {
	tokens := strings.Fields(req.Text)
	if len(tokens) < 2 {
		return "Используйте формат ввода: /untrack <url>"
	}

	url := tokens[1]

	if _, err := h.Service.RemoveLink(req.Context, req.ChatID, url); err != nil {
		return "Не удалось удалить ссылку: " + err.Error()
	}

	return "Ссылка удалена"
}