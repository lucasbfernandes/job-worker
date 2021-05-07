package commands

import (
	"cli/internal/config"
	"errors"
	"flag"
	"os"
)

func (w *WorkerCLI) GetJobs(parameters []string) error {
	getCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := getCmd.String("s", config.GetDefaultServerURL(), "server url")
	apiToken := getCmd.String("t", "", "user api token")

	err := getCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *apiToken == "" {
		return errors.New("api token cannot be empty")
	}

	response, err := w.workerCLIInteractor.GetJobs(*serverURL, *apiToken)
	if err != nil {
		return err
	}

	_, _ = os.Stdout.WriteString(*response + "\n")
	return nil
}
