package dispatcher

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/application/handler"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
)

func TestRegisterAndDispatch(t *testing.T) {
	d := New()

	d.Register("start", "запуск телеграмм бота", handler.Start{})
	d.Register("help", "список доступных команд", handler.Help{
		Descriptions: map[command.Name]string {
			"start": "запуск телеграмм бота",
			"help":  "список доступных команд",
		},
	})

	got, ok := d.Dispatch("start")
	require.True(t, ok)
	require.Equal(t, "Привет! Чтобы посмотреть список доступных команд, воспользуйся командой /help", got)

	require.Len(t, d.Commands(), 2)
}


