package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
)

const (
	createJobPath = "/jobs"
)

func CreateJob(serverURL string, command []string) (*string, error) {
	createJobRequest := dto.NewCreateJobRequest(command)
	createJobResponse, err := requestCreateJob(serverURL, createJobRequest)
	if err != nil {
		return nil, err
	}
	return &createJobResponse.ID, nil
}

func requestCreateJob(serverURL string, createJobRequest *dto.CreateJobRequest) (*dto.CreateJobResponse, error) {
	var createJobResponse dto.CreateJobResponse

	client := resty.New()
	_, err := client.R().
		SetBody(createJobRequest).
		SetResult(&createJobResponse).
		Post(serverURL + createJobPath)

	if err != nil {
		return nil, err
	}

	return &createJobResponse, nil
}
