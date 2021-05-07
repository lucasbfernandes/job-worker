package dto

import (
	jobEntity "server/internal/models/job"
	"time"
)

type GetJobStatusResponse struct {
	Status     string     `json:"status"`
	ExitCode   int        `json:"exitCode"`
	User       string     `json:"user"`
	CreatedAt  *time.Time `json:"createdAt"`
	FinishedAt *time.Time `json:"finishedAt"`
}

func JobStatusResponseFromJob(job *jobEntity.Job) GetJobStatusResponse {
	return GetJobStatusResponse{
		Status:     job.Status,
		ExitCode:   job.ExitCode,
		User:       job.UserID,
		CreatedAt:  job.CreatedAt,
		FinishedAt: job.FinishedAt,
	}
}
