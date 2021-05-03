package job

import (
	"github.com/google/uuid"

	"time"
)

type Job struct {
	ID               string
	Command          []string
	Status           string
	ExitCode         int
	TimeoutInSeconds time.Duration
	CreatedAt        time.Time
	FinishedAt       time.Time
}

// Jobs will start with a -1 ExitCode because this is the default value for
// when processes hasn't exited yet (https://golang.org/pkg/os/#ProcessState.ExitCode)
func NewJob(command []string, timeoutInSeconds time.Duration) Job {
	return Job{
		ID:               uuid.New().String(),
		Command:          command,
		Status:           CREATED,
		ExitCode:         -1,
		TimeoutInSeconds: timeoutInSeconds,
		CreatedAt:        time.Now(),
		FinishedAt:       time.Time{},
	}
}
