package commands

import (
	"errors"
	"fmt"
)

func (w *WorkerCLI) GetJobs(serverURL string) error {
	if serverURL == "" {
		return errors.New("server url cannot be empty")
	}

	formattedJobs, err := w.workerCLIInteractor.GetJobs(serverURL)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *formattedJobs)
	return nil
}
