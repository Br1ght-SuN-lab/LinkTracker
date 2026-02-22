package handlers


func StartHandler() string {
	return "Добро пожаловать! Используйте /help, чтобы посмотреть доступные команды.";
}


func HelpHandler() string {
	return "Доступные команды: \n/start \n/help";
}


func UnknownHandler() string {
	return "Неизвестная команда. Воспользуйтесь командой /help";
}