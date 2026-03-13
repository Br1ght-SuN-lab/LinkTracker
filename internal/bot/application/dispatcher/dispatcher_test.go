package dispatcher

import (
	"testing"
	"context"
	"github.com/stretchr/testify/require"
	"gitlab.education.tbank.ru/backend-academy-go-2025/homeworks/link-tracker/internal/bot/domain/command"
)

type stubHandler struct {
	response string
}

func (h stubHandler) Handle(req command.Request) string {
	return h.response
}
func TestRegisterAndDispatch(t *testing.T) {
	d := New()

	req := command.Request{
		Context: context.Background(),
		ChatID:  123,
		Text:    "/start",
	}

	d.Register("start", "запуск телеграмм бота", stubHandler{
		response: "start response",
	})
	d.Register("help", "список доступных команд", stubHandler{
		response: "help response",
	})

	got, ok := d.Dispatch("start", req)
	require.True(t, ok)
	require.Equal(t, "start response", got)
	require.Len(t, d.Commands(), 2)
}


