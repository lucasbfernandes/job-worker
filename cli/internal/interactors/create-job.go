package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"crypto/tls"
	"errors"
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

	// We are skipping this verification because server has a self-signed certificate
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
		SetBody(createJobRequest).
		SetResult(&createJobResponse).
		SetError(&createJobError).
		Post(serverURL + jobsPath)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		if createJobError.Error != "" {
			return nil, errors.New(createJobError.Error)
		}
		return nil, errors.New("could not create job")
	}

	return &createJobResponse, nil
}
