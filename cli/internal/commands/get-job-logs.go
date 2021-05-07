package commands

import (
	"cli/internal/config"
	"errors"
	"flag"
	"os"
)

func (w *WorkerCLI) GetJobLogs(parameters []string) error {
	logsCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := logsCmd.String("s", config.GetDefaultServerURL(), "server url")
	jobID := logsCmd.String("i", "", "job id")
	apiToken := logsCmd.String("t", "", "user api token")

	err := logsCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *jobID == "" {
		return errors.New("job id cannot be empty")
	}

	if *apiToken == "" {
		return errors.New("api token cannot be empty")
	}

	response, err := w.workerCLIInteractor.GetJobLogs(*serverURL, *jobID, *apiToken)
	if err != nil {
		return err
	}

	_, _ = os.Stdout.WriteString(*response + "\n")
	return nil
}
