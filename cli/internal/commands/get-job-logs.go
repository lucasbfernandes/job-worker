package commands

import (
	"cli/internal/config"
	"errors"
	"flag"
	"fmt"
)

func (w *WorkerCLI) GetJobLogs(parameters []string) error {
	logsCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := logsCmd.String("s", config.GetDefaultServerURL(), "server url")
	jobID := logsCmd.String("i", "", "job id")

	err := logsCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *jobID == "" {
		return errors.New("serverUrl and jobId shouldn't be empty")
	}

	response, err := w.workerCLIInteractor.GetJobLogs(*serverURL, *jobID)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *response)
	return nil
}
