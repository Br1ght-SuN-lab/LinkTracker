package outer

import (
	"fmt"
	"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/dispatcher"
)


func HelpText(d *dispatcher.Dispatcher) string {
	cmds := d.Commands()
	if len(cmds) == 0 {
		return "Команд пока нет."
	}

	var b strings.Builder
	b.WriteString("Доступные команды:\n")
	for _, c := range cmds {
		b.WriteString(fmt.Sprintf("/%s — %s\n", c.Cmd, c.Desc))
	}
	return b.String()
}


func SetMyCommands(bot *tgbotapi.BotAPI, d *dispatcher.Dispatcher) error {
	meta := d.Commands()

	cmds := make([]tgbotapi.BotCommand, 0, len(meta))
	for _, m := range meta {
		// Важно: без "/" — у тебя и так cmd хранится без слэша
		cmds = append(cmds, tgbotapi.BotCommand{
			Command:     m.Cmd,
			Description: m.Desc,
		})
	}

	cfg := tgbotapi.NewSetMyCommands(cmds...)
	_, err := bot.Request(cfg)
	return err
}


func Dispatch(d *dispatcher.Dispatcher, text string) (reply string, ok bool) {
	if text == "" {
		return "", false
	}

	h, exists := d.Find(text)
	if !exists {
		return d.UnknownText(), true
	}
	return h(), true
}