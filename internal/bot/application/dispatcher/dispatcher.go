package dispatcher

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
)

type Handler interface {
	Handle() string
	Name() command.Name
	Description() string 
}

type Dispatcher struct {
	handlers map[command.Name]Handler
}

func New() *Dispatcher { //возвращаем указатель, чтобы иметь возможность менять содержание handlers в будущем
	return &Dispatcher{
		handlers: make(map[command.Name]Handler),
	}
}

func (d *Dispatcher) Register(h Handler) {
	d.handlers[h.Name()] = h
}


func (d *Dispatcher) Dispatch(name command.Name, req command.Request) (string, bool) {
	h, ok := d.handlers[name]
	if !ok {
		return handler.Unknown{}.Handle()
	}

	return h.Handle(req), true
}
