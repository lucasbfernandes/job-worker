package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"cli/internal/security"
	"crypto/tls"
	"errors"
)

func (i *WorkerCLIInteractor) StopJob(serverURL string, jobID string, apiToken string) error {
	err := requestStopJob(serverURL, jobID, apiToken)
	if err != nil {
		return err
	}
	return nil
}

func requestStopJob(serverURL string, jobID string, apiToken string) error {
	var stopJobError dto.JobsError

	bearerToken, err := security.AuthenticateUser(apiToken)
	if err != nil {
		return err
	}

	// We are skipping this verification because server has a self-signed certificate
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
		SetHeader("Authorization", "Bearer "+*bearerToken).
		SetError(&stopJobError).
		Post(serverURL + jobsPath + "/" + jobID + stopJobsPath)

	if err != nil {
		return err
	}

	if response.IsError() {
		if stopJobError.Error != "" {
			return errors.New(stopJobError.Error)
		}
		return errors.New("could not stop job")
	}

	return nil
}
