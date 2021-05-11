package dto

import (
	jobEntity "server/internal/models/job"
	"server/internal/models/user"
)

type CreateJobRequest struct {
	Command []string `json:"command" binding:"required,min=1,dive,min=1"`
}

type CreateJobResponse struct {
	ID string `json:"id"`
}

func (request *CreateJobRequest) ToJob(user *user.User) *jobEntity.Job {
	return jobEntity.NewJob(
		request.Command,
		user.ID,
	)
}
