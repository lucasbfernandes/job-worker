package dto

import "time"

type GetJobStatusResponse struct {
	Status     string     `json:"status"`
	User       string     `json:"user"`
	ExitCode   int        `json:"exitCode"`
	CreatedAt  *time.Time `json:"createdAt"`
	FinishedAt *time.Time `json:"finishedAt"`
}
