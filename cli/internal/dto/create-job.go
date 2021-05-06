package dto

type CreateJobRequest struct {
	Command []string `json:"command"`
}

type CreateJobResponse struct {
	ID string `json:"id"`
}

type CreateJobError struct {
	Error string `json:"error"`
}

func NewCreateJobRequest(command []string) *CreateJobRequest {
	return &CreateJobRequest{
		Command: command,
	}
}
