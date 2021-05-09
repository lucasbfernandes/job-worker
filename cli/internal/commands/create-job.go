package commands

import (
	"cli/internal/config"
	"errors"
	"flag"
	"fmt"
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

	jobID, err := w.workerCLIInteractor.CreateJob(*serverURL, commands)
	if err != nil {
		return err
	}

	fmt.Printf("Created job successfully. Id: %s\n", *jobID)
	return nil
}
