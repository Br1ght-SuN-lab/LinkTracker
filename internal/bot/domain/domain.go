package domain

import(
	"time"
	"context"
)

type Event struct {
	Text    string
	ChatID  int64
	Command string
	Time    time.Time
}

type Name string

const (
	Start Name = "start"
	Help  Name = "help"
	Track Name = "track"
	Cancel Name = "cancel"
	List Name = "list"
	Untrack Name = "untrack"
)

type Request struct {
	Context context.Context
	ChatID  int64
	Text    string
}

type Handler interface {
	Handle(req Request) string
}