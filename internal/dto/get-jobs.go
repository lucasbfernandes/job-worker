package dto

import (
	jobEntity "job-worker/internal/models/job"
	"time"
)

type JobResponse struct {
	ID         string     `json:"id"`
	Command    []string   `json:"command"`
	Status     string     `json:"status"`
	ExitCode   int        `json:"exitCode"`
	CreatedAt  *time.Time `json:"createdAt"`
	FinishedAt *time.Time `json:"finishedAt"`
}

type GetJobsResponse struct {
	Jobs []JobResponse `json:"jobs"`
}

func JobResponseFromJob(job *jobEntity.Job) JobResponse {
	return JobResponse{
		ID:         job.ID,
		Command:    job.Command,
		Status:     job.Status,
		ExitCode:   job.ExitCode,
		CreatedAt:  job.CreatedAt,
		FinishedAt: job.FinishedAt,
	}
}
