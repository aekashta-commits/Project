package todo

import "time"

type Event struct {
	Input     string
	ErrorText string
	CreatedAt time.Time
}
