package commands

import (
	"errors"
	"fmt"
)

func (w *WorkerCLI) CreateJob(serverURL string, commandArray []string) error {
	if serverURL == "" {
		return errors.New("server url cannot be empty")
	}

	if len(commandArray) == 0 {
		return errors.New("exec must receive at least one executable without arguments")
	}

	jobID, err := w.workerCLIInteractor.CreateJob(serverURL, commandArray)
	if err != nil {
		return err
	}

	fmt.Printf("Created job successfully. Id: %s\n", *jobID)
	return nil
}
