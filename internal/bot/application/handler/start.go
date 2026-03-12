package handler

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"

type Start struct{}

func (Start) Name() command.Name {
	return command.Start
}

func (Start) Description() string {
	return "запуск телеграм бота"
}
func (c Start) Handle() string {
	return "Привет! Чтобы посмотреть список доступных команд, воспользуйся командой /help"
}
