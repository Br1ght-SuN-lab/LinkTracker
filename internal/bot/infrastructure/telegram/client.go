package telegram

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/types"
)

type TelegramBot struct {
	api    *tgbotapi.BotAPI
}

func NewTelegramBot(token string) (*TelegramBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot api: %w", err)
	}

	return &TelegramBot{
		api:    api,
	}, nil
}

func (bot *TelegramBot) SetCommands(commands map[command.Name]string) error {
	cmds := make([]tgbotapi.BotCommand, 0, len(commands))
	for k, v := range commands {
		cmds = append(cmds, tgbotapi.BotCommand{
			Command:     string(k),
			Description: v,
		})
	}

	cfg := tgbotapi.NewSetMyCommands(cmds...)
	_, err := bot.api.Request(cfg)
	return err
}

func (bot *TelegramBot) ReceiveMessages(ctx context.Context) <-chan types.Event {
	eventChan := make(chan types.Event, 100)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.api.GetUpdatesChan(u)

	go func() {
		defer close(eventChan) //закрытие моего канала
		defer bot.api.StopReceivingUpdates()

		for {
			select {
			case <-ctx.Done():
				return
			case update, ok := <-updates:
				if !ok {
					return
				}

				event := bot.convertUpdate(&update)
				select {
				case eventChan <- *event:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return eventChan
}

func (bot *TelegramBot) convertUpdate(update *tgbotapi.Update) *types.Event {
	if update.Message != nil {
		msg := update.Message

		event := &types.Event{
			Text:   msg.Text,
			ChatID: msg.Chat.ID,
			Time:   msg.Time(),
		}

		if msg.IsCommand() {
			event.Command = msg.Command()
		}

		return event
	}

	return nil
}

func (bot *TelegramBot) Send(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.api.Send(msg)
	return err
}
