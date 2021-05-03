package dto

import (
	jobEntity "job-worker/internal/models/job"
	"time"
)

type JobResponse struct {
	ID               string
	Command          []string
	Status           string
	ExitCode         int
	TimeoutInSeconds time.Duration
	CreatedAt        time.Time
	FinishedAt       time.Time
}

type GetJobsResponse struct {
	Jobs []JobResponse
}

func JobResponseFromJob(job jobEntity.Job) JobResponse {
	return JobResponse{
		ID:               job.ID,
		Command:          job.Command,
		Status:           job.Status,
		ExitCode:         job.ExitCode,
		TimeoutInSeconds: job.TimeoutInSeconds,
		CreatedAt:        job.CreatedAt,
		FinishedAt:       job.FinishedAt,
	}
}
