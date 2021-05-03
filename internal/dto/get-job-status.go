package dto

import (
	jobEntity "job-worker/internal/models/job"
	"time"
)

type GetJobStatusResponse struct {
	Status     string    `json:"status"`
	ExitCode   int       `json:"exitCode"`
	CreatedAt  time.Time `json:"createdAt"`
	FinishedAt time.Time `json:"finishedAt"`
}

func JobStatusResponseFromJob(job jobEntity.Job) GetJobStatusResponse {
	return GetJobStatusResponse{
		Status:     job.Status,
		ExitCode:   job.ExitCode,
		CreatedAt:  job.CreatedAt,
		FinishedAt: job.FinishedAt,
	}
}
