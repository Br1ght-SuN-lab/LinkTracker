package command

type Meta struct{
	Cmd string
	Desc string
}

type Provider interface {
	Commands() []Meta
}