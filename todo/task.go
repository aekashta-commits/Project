package todo

import "time"

type Task struct {
	Title       string
	Description string
	CreatedAt   time.Time
	IsDone      bool
	DoneAt      *time.Time
}
