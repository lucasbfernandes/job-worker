package dto

import (
	jobEntity "job-worker/internal/models/job"
	"time"
)

type GetJobStatusResponse struct {
	Status     string
	ExitCode   int
	CreatedAt  time.Time
	FinishedAt time.Time
}

func JobStatusResponseFromJob(job jobEntity.Job) GetJobStatusResponse {
	return GetJobStatusResponse{
		Status:     job.Status,
		ExitCode:   job.ExitCode,
		CreatedAt:  job.CreatedAt,
		FinishedAt: job.FinishedAt,
	}
}
