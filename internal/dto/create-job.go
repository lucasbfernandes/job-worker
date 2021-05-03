package dto

import (
	jobEntity "job-worker/internal/models/job"
	"time"
)

type CreateJobRequest struct {
	Command          []string
	TimeoutInSeconds int
}

type CreateJobResponse struct {
	ID string
}

func (request *CreateJobRequest) ToJob() jobEntity.Job {
	return jobEntity.NewJob(
		request.Command,
		time.Duration(request.TimeoutInSeconds),
	)
}
