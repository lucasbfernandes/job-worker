package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"errors"
	"fmt"
)

const (
	getJobsPath = "/jobs"
)

func (i *WorkerCLIInteractor) GetJobs(serverURL string) (*string, error) {
	getJobsResponse, err := requestGetJobs(serverURL)
	if err != nil {
		return nil, err
	}

	noJobsResponse := "No jobs found."
	if len(getJobsResponse.Jobs) == 0 {
		return &noJobsResponse, nil
	}

	parsedResponse := parseGetJobsResponse(getJobsResponse)
	return parsedResponse, nil
}

func requestGetJobs(serverURL string) (*dto.GetJobsResponse, error) {
	var getJobsResponse dto.GetJobsResponse
	var getJobsError dto.JobsError

	client := resty.New()
	response, err := client.R().
		SetResult(&getJobsResponse).
		SetError(&getJobsError).
		Get(serverURL + getJobsPath)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		if getJobsError.Error != "" {
			return nil, errors.New(getJobsError.Error)
		}
		return nil, errors.New("could not get jobs")
	}

	return &getJobsResponse, nil
}

func parseGetJobsResponse(getJobsResponse *dto.GetJobsResponse) *string {
	parsedResponse := ""
	for index, job := range getJobsResponse.Jobs {
		parsedResponse += fmt.Sprintf(
			"%d\nid: %s\ncommand: %s\nstatus: %s\ncreatedAt: %s\nfinishedAt: %s\n",
			index+1, job.ID, job.Command, job.Status, job.CreatedAt, job.FinishedAt,
		)
	}
	return &parsedResponse
}
