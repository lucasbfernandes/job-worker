package dto

import "time"

type GetJobStatusResponse struct {
	Status     string     `json:"status"`
	ExitCode   int        `json:"exitCode"`
	CreatedAt  *time.Time `json:"createdAt"`
	FinishedAt *time.Time `json:"finishedAt"`
}
