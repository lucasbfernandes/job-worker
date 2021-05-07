package job

import (
	"github.com/google/uuid"

	"server/pkg/worker"
	"time"
)

type Job struct {
	ID         string
	Command    []string
	Status     string
	ExitCode   int
	CreatedAt  *time.Time
	FinishedAt *time.Time
	UserID     string

	process *worker.Process
}

// Jobs will start with a -1 ExitCode because this is the default value for
// when processes hasn't exited yet (https://golang.org/pkg/os/#ProcessState.ExitCode)
func NewJob(command []string, userID string) *Job {
	createdAt := time.Now()
	return &Job{
		ID:         uuid.New().String(),
		Command:    command,
		Status:     CREATED,
		ExitCode:   -1,
		CreatedAt:  &createdAt,
		FinishedAt: nil,
		UserID:     userID,

		process: nil,
	}
}

func (job *Job) SetProcess(process *worker.Process) {
	job.process = process
}

func (job *Job) GetProcess() *worker.Process {
	return job.process
}
