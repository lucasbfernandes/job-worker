package commands

import (
	"cli/internal/config"
	"errors"
	"flag"
	"fmt"
)

func (w *WorkerCLI) GetJobStatus(parameters []string) error {
	statusCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := statusCmd.String("s", config.GetDefaultServerURL(), "server url")
	jobID := statusCmd.String("i", "", "job id")

	err := statusCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *jobID == "" {
		return errors.New("job id cannot be empty")
	}

	formattedStatus, err := w.workerCLIInteractor.GetJobStatus(*serverURL, *jobID)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *formattedStatus)
	return nil
}
