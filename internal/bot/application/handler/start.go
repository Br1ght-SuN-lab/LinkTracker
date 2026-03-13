package handler

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
)

type Start struct {
	Service *application.Service
}

func (s Start) Handle(req command.Request) string {
	if err := s.Service.RegisterChat(req.Context, req.ChatID); err != nil {
		return "Ошибка регистрации чата"
	}
	return "Привет! Чтобы посмотреть список доступных команд, воспользуйся командой /help"
}