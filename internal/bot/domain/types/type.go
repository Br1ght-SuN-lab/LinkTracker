package types

import "time"

type LogLevel string

const (
	LogDebug LogLevel = "debug"
	LogInfo  LogLevel = "info"
	LogWarn  LogLevel = "warn"
	LogError LogLevel = "error"
)

type Event struct {
	Text    string
	ChatID  int64
	Command string
	Time    time.Time
}
