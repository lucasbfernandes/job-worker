package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"cli/internal/security"
	"crypto/tls"
	"errors"
	"fmt"
)

func (i *WorkerCLIInteractor) GetJobStatus(serverURL string, jobID string, apiToken string) (*string, error) {
	getJobStatusResponse, err := requestGetJobStatus(serverURL, jobID, apiToken)
	if err != nil {
		return nil, err
	}

	parsedResponse := parseGetJobStatusResponse(getJobStatusResponse)
	return parsedResponse, nil
}

func requestGetJobStatus(serverURL string, jobID string, apiToken string) (*dto.GetJobStatusResponse, error) {
	var getJobStatusResponse dto.GetJobStatusResponse
	var getJobStatusError dto.JobsError

	bearerToken, err := security.AuthenticateUser(apiToken)
	if err != nil {
		return nil, err
	}

	// We are skipping this verification because server has a self-signed certificate
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
		SetHeader("Authorization", "Bearer "+*bearerToken).
		SetResult(&getJobStatusResponse).
		SetError(&getJobStatusError).
		Get(serverURL + jobsPath + "/" + jobID + getJobStatusPath)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		if response.StatusCode() == 401 {
			return nil, errors.New("failed authentication - unauthorized")
		}
		if getJobStatusError.Error != "" {
			return nil, errors.New(getJobStatusError.Error)
		}
		return nil, errors.New("could not get job status")
	}

	return &getJobStatusResponse, nil
}

func parseGetJobStatusResponse(response *dto.GetJobStatusResponse) *string {
	parsedResponse := fmt.Sprintf(
		"\nstatus: %s\nuser: %s\ncreatedAt: %s\nfinishedAt: %s\nexitCode: %d\n",
		response.Status, response.User, response.CreatedAt.Format(dateLayout), response.FinishedAt.Format(dateLayout), response.ExitCode,
	)
	return &parsedResponse
}
