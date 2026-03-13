package command

import "context"

type Name string

const (
	Start Name = "start"
	Help  Name = "help"
)

type Request struct {
	Context context.Context
	ChatID  int64
	Text    string
}
type Handler interface {
	Handle(req Request) string
}

type Meta struct {
	Name Name
	Desc string
}