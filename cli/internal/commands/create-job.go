package commands

import (
	"cli/internal/config"
	"errors"
	"flag"
	"fmt"
	"os"
)

func (w *WorkerCLI) CreateJob(parameters []string) error {
	execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := execCmd.String("s", config.GetDefaultServerURL(), "server url")
	apiToken := execCmd.String("t", "", "user api token")

	err := execCmd.Parse(parameters)
	if err != nil {
		return errors.New("failed to parse exec command line arguments")
	}

	if *apiToken == "" {
		return errors.New("api token cannot be empty")
	}

	commands := execCmd.Args()
	if len(commands) == 0 {
		return errors.New("exec must receive at least one executable without arguments")
	}

	response, err := w.workerCLIInteractor.CreateJob(*serverURL, commands, *apiToken)
	if err != nil {
		return err
	}

	_, _ = os.Stdout.WriteString(fmt.Sprintf("Created job successfully. Id: %s.\n", *response))
	return nil
}
