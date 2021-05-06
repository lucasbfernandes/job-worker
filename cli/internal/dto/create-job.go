package dto

type CreateJobRequest struct {
	Command []string `json:"command"`
}

type CreateJobResponse struct {
	ID string `json:"id"`
}

func NewCreateJobRequest(command []string) *CreateJobRequest {
	return &CreateJobRequest{
		Command: command,
	}
}
