package handler

import (
	"log/slog"
	"strconv"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain"
)

type List struct {
	Service *application.Service
	Logger  *slog.Logger
}

func (h List) Handle(req domain.Request) string {
	links, err := h.Service.ListLinks(req.Context, req.ChatID)
	if err != nil {
		h.Logger.Info("failed to get links", "error", err)
		return "Не удалось получить список ссылок"
	}

	if len(links) == 0 {
		return "Список отслеживаемых ссылок пуст"
	}

	var b strings.Builder
	b.WriteString("Отслеживаемые ссылки:\n")

	for i, link := range links {
		b.WriteString(strconv.Itoa(i + 1))
		b.WriteString(": ")
		b.WriteString(link.URL)
		if i != len(links)-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}
