package dispatcher

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
)

type Handler interface {
	Handle(command.Request) string
}

type Dispatcher struct {
	handlers     map[command.Name]Handler
	descriptions map[command.Name]string
}

func New() *Dispatcher {
	return &Dispatcher{
		handlers:     make(map[command.Name]Handler),
		descriptions: make(map[command.Name]string),
	}
}

func (d *Dispatcher) Register(name command.Name, desc string, h Handler) {
	d.handlers[name] = h
	d.descriptions[name] = desc
}

func (d *Dispatcher) Dispatch(name command.Name, req command.Request) string {
	h, ok := d.handlers[name]
	if !ok {
		return handler.Unknown{}.Handle()
	}

	return h.Handle(req)
}