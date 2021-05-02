package dto

import (
	"job-worker/internal/models/job"
	"time"
)

type CreateJobRequest struct {
	Command          []string
	TimeoutInSeconds int
}

type CreateJobResponse struct {
	ID string
}

func (request *CreateJobRequest) ToJob() job.Job {
	return job.NewJob(
		request.Command,
		time.Duration(request.TimeoutInSeconds),
	)
}
