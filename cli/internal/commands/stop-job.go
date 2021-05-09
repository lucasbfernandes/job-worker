package commands

import (
	"cli/internal/config"
	"errors"
	"flag"
	"fmt"
)

func (w *WorkerCLI) StopJob(parameters []string) error {
	stopCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := stopCmd.String("s", config.GetDefaultServerURL(), "server url")
	jobID := stopCmd.String("i", "", "job id")

	err := stopCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *jobID == "" {
		return errors.New("job id cannot be empty")
	}

	err = w.workerCLIInteractor.StopJob(*serverURL, *jobID)
	if err != nil {
		return err
	}

	fmt.Printf("Job stopped successfully.\n")
	return nil
}
