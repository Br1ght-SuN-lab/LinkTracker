package handlers

func Start() string {
	return "Добро пожаловать! Используйте /help, чтобы посмотреть доступные команды."
}

func Help(getHelpText func() string) (func() string) {
	return func() string {
		return getHelpText()
	}
}

func Unknown() string {
	return "Неизвестная команда. Воспользуйтесь командой /help"
}