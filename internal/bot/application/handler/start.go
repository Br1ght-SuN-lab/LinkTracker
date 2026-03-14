package handler

import (
	"log/slog"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain"
)

type Start struct {
	Logger  *slog.Logger
	Service *application.Service
}

func (s Start) Handle(req domain.Request) string {
	if err := s.Service.RegisterChat(req.Context, req.ChatID); err != nil {
		s.Logger.Info("faied on register chat", "error", err)
		return "Ошибка регистрации чата"
	}
	return "Привет! Чтобы посмотреть список доступных команд, воспользуйся командой /help"
}