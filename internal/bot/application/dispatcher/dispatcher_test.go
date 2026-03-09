package dispatcher

import (
	"testing"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler"
	"github.com/stretchr/testify/require"
)

func TestRegisterAndDispatch(t *testing.T) {
	d := New()

	d.Register("start", "запуск телеграмм бота", handler.Start)

	got, ok := d.Dispatch("start")
	require.True(t, ok)
	require.Equal(t, "Добро пожаловать! Используйте /help, чтобы посмотреть доступные команды.", got)
}


func TestCommands(t *testing.T) {
	d := New()

	d.Register("start", "запуск телеграмм бота", handler.Start)
	d.Register("help", "список команд", handler.Help(d))

	cmds := d.Commands()
	require.Len(t, cmds, 2)
}
