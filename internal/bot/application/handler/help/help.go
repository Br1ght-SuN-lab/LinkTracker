package help

import (
	"fmt"
	"strings"
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


func Help(getHelpText func() string) (func() string) {
	return func() string {
		return getHelpText()
	}
}