package dispatcher

import (
	"fmt"
	"sort"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HandlerFunc func() string //тип для фукнции

type Command struct {
	Handler HandlerFunc
	Desc string 
}
type Dispatcher struct {
	handlers map[string]Command
	unknown string 
}

func New() *Dispatcher { //возвращаем указатель, чтобы иметь возможность менять содержание handlers в будущем
	return &Dispatcher{
		handlers: map[string]Command{},
		unknown: "Неизвестная команда. Воспользуйтесь командой /help",
	}
}

func (d *Dispatcher) Register(cmd string, desc string, h HandlerFunc) {
	d.handlers[cmd] = Command{
        Handler: h,
        Desc: desc,
    }
}


func (d *Dispatcher) Dispatch(text string) (reply string, flag bool) {
	if text == "" {
		return "", false
	}

	h, exists := d.handlers[text]
	if !exists {
		return d.unknown, true
	}

	return h.Handler(), true
}


func (d *Dispatcher) HelpText() string {
    if len(d.handlers) == 0 {
        return "Команд пока нет."
    }

    keys := make([]string, 0, len(d.handlers))
    for k := range d.handlers {
        keys = append(keys, k)
    }

    var res strings.Builder
    res.WriteString("Доступные команды:\n")
    for _, k := range keys {
        c := d.handlers[k]
        res.WriteString(fmt.Sprintf("/%s — %s\n", k, c.Desc))
    }
    return res.String()
}


func (d *Dispatcher) SetMyCommands(bot *tgbotapi.BotAPI) error {
	//мне кажется прикольно отсортировать команды
	keys := make([]string, 0, len(d.handlers))
	for k := range d.handlers {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	cmds := make([]tgbotapi.BotCommand, 0, len(d.handlers))
	for _, key := range keys {
		cmds = append(cmds, tgbotapi.BotCommand{
			Command: key,
			Description: d.handlers[key].Desc,
		})
	}

	cfg := tgbotapi.NewSetMyCommands(cmds...)

	_, err := bot.Request(cfg)
	return err
}