package handler

type Start struct{}

func (c Start) Handle() string {
	return "Привет! Чтобы посмотреть список доступных команд, воспользуйся командой /help"
}
