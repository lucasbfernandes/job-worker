package dto

import (
	"github.com/google/uuid"
	"job-worker/internal/models/job"
	"time"
)

type CreateJobRequest struct {
	Command          []string
	TimeoutInSeconds int
}

type CreateJobResponse struct {
	ID uuid.UUID
}

func (request *CreateJobRequest) ToJob() job.Job {
	return job.NewJob(
		request.Command,
		time.Duration(request.TimeoutInSeconds),
	)
}
