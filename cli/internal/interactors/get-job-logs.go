package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"crypto/tls"
	"errors"
)

func (i *WorkerCLIInteractor) GetJobLogs(serverURL string, jobID string) (*string, error) {
	getJobLogsResponse, err := requestGetJobLogs(serverURL, jobID)
	if err != nil {
		return nil, err
	}
	return getJobLogsResponse, nil
}

func requestGetJobLogs(serverURL string, jobID string) (*string, error) {
	var getJobLogsError dto.JobsError

	// We are skipping this verification because server has a self-signed certificate
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
		SetError(&getJobLogsError).
		Get(serverURL + jobsPath + "/" + jobID + getJobLogsPath)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		if getJobLogsError.Error != "" {
			return nil, errors.New(getJobLogsError.Error)
		}
		return nil, errors.New("could not get job logs")
	}

	logs := response.String()
	return &logs, nil
}
