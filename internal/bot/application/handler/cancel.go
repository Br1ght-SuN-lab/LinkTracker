package handler

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain"
)

type Cancel struct {
	Service *application.Service
}

func (h Cancel) Handle(req domain.Request) string {
	if !h.Service.HasActiveTrackDialog(req.ChatID) {
		return "Нет активного диалога"
	}

	h.Service.DeleteTrackDialog(req.ChatID)
	return "Диалог /track отменен"
}