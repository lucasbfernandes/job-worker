package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"crypto/tls"
	"errors"
	"fmt"
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

	// We are skipping this verification because server has a self-signed certificate
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
		SetResult(&getJobsResponse).
		SetError(&getJobsError).
		Get(serverURL + jobsPath)

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
			"\n%d\nid: %s\ncommand: %s\nstatus: %s\ncreatedAt: %s\nfinishedAt: %s\n",
			index+1, job.ID, job.Command, job.Status, job.CreatedAt.Format(dateLayout), job.FinishedAt.Format(dateLayout),
		)
	}
	return &parsedResponse
}
