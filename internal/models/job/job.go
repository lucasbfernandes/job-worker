package job

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	Id uuid.UUID
	Command []string
	Status string
	TimeoutInSeconds time.Duration
	CreatedAt time.Time
	FinishedAt time.Time
}

func NewJob(command []string, timeoutInSeconds time.Duration) Job {
	return Job{
		Id:         uuid.New(),
		Command:    command,
		Status:     CREATED,
		TimeoutInSeconds: timeoutInSeconds,
		CreatedAt:  time.Now(),
		FinishedAt: time.Time{},
	}
}
