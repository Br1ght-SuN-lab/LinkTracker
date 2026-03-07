package dispatcher

import "gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler/unknown"

type HandlerFunc func() string //тип для фукнции

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

// найти команду, нужно в infrastructure/outer/
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

type CommandMeta struct {
	Cmd  string
	Desc string
}


func (d *Dispatcher) Commands() []CommandMeta {
	out := make([]CommandMeta, 0, len(d.handlers))
	for k, v := range d.handlers {
		out = append(out, CommandMeta{
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