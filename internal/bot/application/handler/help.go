package handler

import (
	"fmt"
	"sort"
	"strings"

	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
)

type Help struct {
	Descriptions map[command.Name]string
}

func (c Help) Handle() string {
	if len(c.Descriptions) == 0 {
		return "Команд пока нет."
	}

	var names []string
	for name := range c.Descriptions {
		names = append(names, string(name))
	}
	sort.Strings(names)
	var b strings.Builder
	b.WriteString("Доступные команды:\n")
	for _, name := range names {
		b.WriteString(fmt.Sprintf("/%s — %s\n", name, c.Descriptions[command.Name(name)]))
	}

	return b.String()
}
