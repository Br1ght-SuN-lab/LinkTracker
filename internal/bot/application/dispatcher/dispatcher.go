package dispatcher

import (
	"strings"
	"fmt"
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