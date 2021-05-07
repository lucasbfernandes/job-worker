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
	apiToken := statusCmd.String("t", "", "user api token")

	err := statusCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *jobID == "" {
		return errors.New("job id cannot be empty")
	}

	if *apiToken == "" {
		return errors.New("api token cannot be empty")
	}

	response, err := w.workerCLIInteractor.GetJobStatus(*serverURL, *jobID, *apiToken)
	if err != nil {
		return err
	}

	_, _ = os.Stdout.WriteString(*response + "\n")
	return nil
}
