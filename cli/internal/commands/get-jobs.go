package commands

import (
	"errors"
	"fmt"
)

func (w *WorkerCLI) GetJobs(serverURL string, apiToken string) error {
	if serverURL == "" {
		return errors.New("server url cannot be empty")
	}

	if apiToken == "" {
		return errors.New("api token cannot be empty")
	}

	formattedJobs, err := w.workerCLIInteractor.GetJobs(serverURL, apiToken)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *formattedJobs)
	return nil
}
