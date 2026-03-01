package dispatcher

import (
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handlers"
)

type HandlerFunc func() string //тип для фукнции

type Dispatcher struct {
	handlers map[string]HandlerFunc
	unknown HandlerFunc
}

func New() *Dispatcher { //возвращаем указатель, чтобы иметь возможность менять содержание handlers в будущем
	return &Dispatcher{
		handlers: map[string]HandlerFunc{
			"/start": handlers.StartHandler,
			"/help":  handlers.HelpHandler,
		},
		unknown: handlers.UnknownHandler,
	}
}

func (d *Dispatcher) Register(cmd string, h HandlerFunc) {
	d.handlers[cmd] = h //вместо h любой обработчик
}

func (d *Dispatcher) Dispatch(text string) (reply string, flag bool) {
	if text == "" {
		return "", false
	}

	h, exists := d.handlers[text]
	if !exists {
		return d.unknown(), true
	}

	return h(), true
}
