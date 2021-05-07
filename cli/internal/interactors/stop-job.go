package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"crypto/tls"
	"errors"
)

func (i *WorkerCLIInteractor) StopJob(serverURL string, jobID string) error {
	err := requestStopJob(serverURL, jobID)
	if err != nil {
		return err
	}
	return nil
}

func requestStopJob(serverURL string, jobID string) error {
	var stopJobError dto.JobsError

	// We are skipping this verification because server has a self-signed certificate
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	response, err := client.R().
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
