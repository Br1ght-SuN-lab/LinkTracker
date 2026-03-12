package handler

type Unknown struct{}

func (c Unknown) Handle() string {
	return "Неизвестная команда. Воспользуйтесь командой /help"
}
