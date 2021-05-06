package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"errors"
)

const (
	jobsPath     = "/jobs/"
	stopJobsVerb = "/stop"
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

	client := resty.New()
	response, err := client.R().
		SetError(&stopJobError).
		Post(serverURL + jobsPath + jobID + stopJobsVerb)

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
