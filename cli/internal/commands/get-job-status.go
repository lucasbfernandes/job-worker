package commands

import (
	"errors"
	"fmt"
)

func (w *WorkerCLI) GetJobStatus(serverURL string, jobID string) error {
	if serverURL == "" {
		return errors.New("server url cannot be empty")
	}

	if jobID == "" {
		return errors.New("job id cannot be empty")
	}

	formattedStatus, err := w.workerCLIInteractor.GetJobStatus(serverURL, jobID)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *formattedStatus)
	return nil
}
