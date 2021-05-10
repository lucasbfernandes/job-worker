package commands

import (
	"errors"
	"fmt"
)

func (w *WorkerCLI) GetJobLogs(serverURL string, jobID string, apiToken string) error {
	if serverURL == "" {
		return errors.New("server url cannot be empty")
	}

	if jobID == "" {
		return errors.New("job id cannot be empty")
	}

	if apiToken == "" {
		return errors.New("api token cannot be empty")
	}

	formattedLogs, err := w.workerCLIInteractor.GetJobLogs(serverURL, jobID, apiToken)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *formattedLogs)
	return nil
}
