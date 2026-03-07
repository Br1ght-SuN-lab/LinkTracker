package telegram

import "time"

type MyMessage struct {
    Text      string
    ChatID    int64
    Command   string
    Time      time.Time
}

type MyEvent struct {
	Type     string 
	Message *MyMessage
}

