package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"errors"
)

const (
	createJobPath = "/jobs"
)

func (i *WorkerCLIInteractor) CreateJob(serverURL string, command []string) (*string, error) {
	createJobRequest := dto.NewCreateJobRequest(command)
	createJobResponse, err := requestCreateJob(serverURL, createJobRequest)
	if err != nil {
		return nil, err
	}
	return &createJobResponse.ID, nil
}

func requestCreateJob(serverURL string, createJobRequest *dto.CreateJobRequest) (*dto.CreateJobResponse, error) {
	var createJobResponse dto.CreateJobResponse
	var createJobError dto.JobsError

	client := resty.New()
	response, err := client.R().
		SetBody(createJobRequest).
		SetResult(&createJobResponse).
		SetError(&createJobError).
		Post(serverURL + createJobPath)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		return nil, errors.New(createJobError.Error)
	}

	return &createJobResponse, nil
}
