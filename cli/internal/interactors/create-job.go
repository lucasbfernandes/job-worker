package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"cli/internal/security"
	"crypto/tls"
	"errors"
)

func (i *WorkerCLIInteractor) CreateJob(serverURL string, command []string, apiToken string) (*string, error) {
	createJobRequest := dto.NewCreateJobRequest(command)
	createJobResponse, err := requestCreateJob(serverURL, createJobRequest, apiToken)
	if err != nil {
		return nil, err
	}
	return &createJobResponse.ID, nil
}

func requestCreateJob(serverURL string, createJobRequest *dto.CreateJobRequest, apiToken string) (*dto.CreateJobResponse, error) {
	var createJobResponse dto.CreateJobResponse
	var createJobError dto.JobsError

	bearerToken, err := security.AuthenticateUser(apiToken)
	if err != nil {
		return nil, err
	}

	// We are skipping this verification because server has a self-signed certificate
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
		SetBody(createJobRequest).
		SetHeader("Authorization", "Bearer "+*bearerToken).
		SetResult(&createJobResponse).
		SetError(&createJobError).
		Post(serverURL + jobsPath)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		if response.StatusCode() == 401 {
			return nil, errors.New("failed authentication - unauthorized")
		}
		if createJobError.Error != "" {
			return nil, errors.New(createJobError.Error)
		}
		return nil, errors.New("could not create job")
	}

	return &createJobResponse, nil
}
