package telegram

import (
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/dispatcher"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler/help"
    "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler/start"
)

type TelegramBot struct {
	Api tgbotapi.BotAPI
	logger slog.Logger
}


func NewTelegramBot(token string, logger slog.Logger) (*TelegramBot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot api: %w", err)
	}

	return &TelegramBot{
		Api:    *api,
		logger: logger,
	}, nil
}


func (bot *TelegramBot) RegistrationCommands(d *dispatcher.Dispatcher) {
    d.Register("start", "запуск телеграмм бота", handler.Start)
	getHelpText := func() string {return help.HelpText(d)}
	d.Register("help", "список доступных команд", help.Help(func() string {return getHelpText()}))
}


func (bot *TelegramBot) SetMyCommands(d *dispatcher.Dispatcher) error {
	meta := d.Commands()

	cmds := make([]tgbotapi.BotCommand, 0, len(meta))
	for _, m := range meta {
		cmds = append(cmds, tgbotapi.BotCommand{
			Command:     m.Cmd,
			Description: m.Desc,
		})
	}

	cfg := tgbotapi.NewSetMyCommands(cmds...)
	_, err := bot.Api.Request(cfg)
	return err
}


func (bot *TelegramBot) ReceiveMessages() <-chan MyEvent {
    eventChan := make(chan MyEvent, 100)
    
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates := bot.Api.GetUpdatesChan(u)
    
    go func() {
        for update := range updates {
            event := bot.convertUpdate(&update)
            eventChan <- event
        }
    }()
    
    return eventChan
}


func (bot *TelegramBot) convertUpdate(update *tgbotapi.Update) MyEvent {
	if update.Message != nil {
		msg := update.Message
		
		myMsg := &MyMessage{
			Text: msg.Text,
			ChatID: msg.Chat.ID,
			Time: msg.Time(),
		}

		if msg.IsCommand() {
			myMsg.Command = msg.Command()
		}

		eventType := "message"
		if msg.IsCommand() {
			eventType = "command"
		}

		return MyEvent{
			Type: eventType,
			Message: myMsg,
		}
	}

	return MyEvent{Type: "unknown"}
}


func (b *TelegramBot) Send(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.Api.Send(msg)
	return err
}