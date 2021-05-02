package dto

type CreateJobRequest struct {
	Command          []string
	TimeoutInSeconds int
}

type CreateJobResponse struct {
	ID string
}
