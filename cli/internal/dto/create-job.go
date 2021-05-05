package dto

type CreateJobRequest struct {
	command []string
}

type CreateJobResponse struct {
	ID string
}

func NewCreateJobRequest(command []string) *CreateJobRequest {
	return &CreateJobRequest{
		command: command,
	}
}
