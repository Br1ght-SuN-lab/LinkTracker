package types

import "time"

type Event struct {
    Text      string
    ChatID    int64
    Command   string
    Time      time.Time
}

