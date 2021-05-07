package dto

import "time"

type JobResponse struct {
	ID         string     `json:"id"`
	Command    []string   `json:"command"`
	User       string     `json:"user"`
	Status     string     `json:"status"`
	ExitCode   int        `json:"exitCode"`
	CreatedAt  *time.Time `json:"createdAt"`
	FinishedAt *time.Time `json:"finishedAt"`
}

type GetJobsResponse struct {
	Jobs []JobResponse `json:"jobs"`
}
