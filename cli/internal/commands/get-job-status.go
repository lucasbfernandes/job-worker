package commands

import (
	"cli/internal/config"
	"errors"
	"flag"
	"os"
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
		return errors.New("jobID shouldn't be empty")
	}

	response, err := w.workerCLIInteractor.GetJobStatus(*serverURL, *jobID)
	if err != nil {
		return err
	}

	_, _ = os.Stdout.WriteString(*response + "\n")
	return nil
}
