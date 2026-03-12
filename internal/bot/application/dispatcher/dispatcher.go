package dispatcher

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
)

type Dispatcher struct {
	handlers map[command.Name]command.Handler
	meta     map[command.Name]string
}

func New() *Dispatcher { //возвращаем указатель, чтобы иметь возможность менять содержание handlers в будущем
	return &Dispatcher{
		handlers: make(map[command.Name]command.Handler),
		meta:     make(map[command.Name]string),
	}
}

func (d *Dispatcher) Register(name command.Name, desc string, h command.Handler) {
	d.handlers[name] = h
	d.meta[name] = desc
}

func (d *Dispatcher) Commands() []command.Meta {
	out := make([]command.Meta, 0, len(d.meta))
	for name, desc := range d.meta {
		out = append(out, command.Meta{
			Name: name,
			Desc: desc,
		})
	}
	return out
}

func (d *Dispatcher) Dispatch(name command.Name) (string, bool) {
	h, ok := d.handlers[name]
	if !ok {
		return "", false
	}
	return h.Handle(), true
}
