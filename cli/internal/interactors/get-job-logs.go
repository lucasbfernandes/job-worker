package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"cli/internal/security"
	"crypto/tls"
	"errors"
)

func (i *WorkerCLIInteractor) GetJobLogs(serverURL string, jobID string, apiToken string) (*string, error) {
	getJobLogsResponse, err := requestGetJobLogs(serverURL, jobID, apiToken)
	if err != nil {
		return nil, err
	}
	return getJobLogsResponse, nil
}

func requestGetJobLogs(serverURL string, jobID string, apiToken string) (*string, error) {
	var getJobLogsError dto.JobsError

	bearerToken, err := security.AuthenticateUser(apiToken)
	if err != nil {
		return nil, err
	}

	// We are skipping this verification because server has a self-signed certificate
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
		SetHeader("Authorization", "Bearer "+*bearerToken).
		SetError(&getJobLogsError).
		Get(serverURL + jobsPath + "/" + jobID + getJobLogsPath)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		if response.StatusCode() == 401 {
			return nil, errors.New("failed authentication - unauthorized")
		}
		if getJobLogsError.Error != "" {
			return nil, errors.New(getJobLogsError.Error)
		}
		return nil, errors.New("could not get job logs")
	}

	logs := response.String()
	return &logs, nil
}
