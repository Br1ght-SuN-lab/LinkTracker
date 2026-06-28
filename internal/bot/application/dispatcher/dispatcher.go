package dispatcher

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain"
)

type Handler interface {
	Handle(domain.Request) string
}

type Dispatcher struct {
	handlers     map[domain.Name]Handler
	descriptions map[domain.Name]string
}

func New() *Dispatcher {
	return &Dispatcher{
		handlers:     make(map[domain.Name]Handler),
		descriptions: make(map[domain.Name]string),
	}
}

func (d *Dispatcher) Register(name domain.Name, desc string, h Handler) {
	d.handlers[name] = h
	d.descriptions[name] = desc
}

func (d *Dispatcher) Dispatch(name domain.Name, req domain.Request) string {
	h, ok := d.handlers[name]
	if !ok {
		return handler.Unknown{}.Handle()
	}

	return h.Handle(req)
}