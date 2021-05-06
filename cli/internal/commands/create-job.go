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

	err := execCmd.Parse(parameters)
	if err != nil {
		return errors.New("failed to parse exec command line arguments")
	}

	commands := execCmd.Args()
	if len(commands) == 0 {
		return errors.New("exec must receive at least one executable without arguments")
	}

	response, err := w.workerCLIInteractor.CreateJob(*serverURL, commands)
	if err != nil {
		return err
	}

	_, _ = os.Stdout.WriteString(fmt.Sprintf("Created job successfully. Id: %s.\n", *response))
	return nil
}
