package command

type Name string

const (
	Start Name = "start"
	Help  Name = "help"
)

type Handler interface {
	Handle() string
}

type Meta struct {
	Name Name
	Desc string
}
