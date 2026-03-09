package dispatcher

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
)

type HandlerFunc func() string //тип для функции

type Command struct {
	Handler HandlerFunc
	Desc    string
}
type Dispatcher struct {
	handlers map[string]Command
	unknown HandlerFunc
}


func New() *Dispatcher { //возвращаем указатель, чтобы иметь возможность менять содержание handlers в будущем
	return &Dispatcher{
		handlers: map[string]Command{},
		unknown:  handler.Unknown,
	}
}

func (d *Dispatcher) Register(cmd string, desc string, h HandlerFunc) {
	d.handlers[cmd] = Command{
		Handler: h,
		Desc:    desc,
	}
}


func (d *Dispatcher) RegistrationCommands() {
    d.Register("start", "запуск телеграмм бота", handler.Start)
	d.Register("help", "список доступных команд", handler.Help(d))
}


func (d *Dispatcher) Find(cmd string) (HandlerFunc, bool) {
	c, ok := d.handlers[cmd]
	if !ok {
		return nil, false
	}

	return c.Handler, true
}


func (d *Dispatcher) UnknownText() string {
	return d.unknown()
}


func (d *Dispatcher) Commands() []command.Meta {
	out := make([]command.Meta, 0, len(d.handlers))
	for k, v := range d.handlers {
		out = append(out, command.Meta{
			Cmd: k, 
			Desc: v.Desc,
		})
	}
	return out
}


func Dispatch(d *Dispatcher, text string) (reply string, ok bool) {
	if text == "" {
		return "", false
	}

	h, exists := d.Find(text)
	if !exists {
		return d.UnknownText(), true
	}
	return h(), true
}