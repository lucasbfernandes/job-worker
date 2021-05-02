package models

import "time"

// TODO rethink finishedAt. It might be null sometimes
type Job struct {
	Id string
	Command []string
	Status string
	CreatedAt time.Time
	FinishedAt time.Time
}
