package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func RegisterCommands(bot *tgbotapi.BotAPI) error {
	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Запуск бота"},
		{Command: "help", Description: "Список команд"},
	}

	cfg := tgbotapi.NewSetMyCommands(commands...)
	_, err := bot.Request(cfg)
	return err
}