package commands

import (
	"cli/internal/config"
	"cli/internal/interactors"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func (w *WorkerCLI) CreateJob(parameters []string) error {
	// TODO use flag.ContinueOnError?
	execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := execCmd.String("s", config.GetDefaultServerURL(), "server url")
	executable := execCmd.String("c", "", "command to be executed on the server")

	err := execCmd.Parse(parameters)
	if err != nil {
		return errors.New("failed to parse exec command line arguments")
	}

	if *executable == "" {
		return errors.New("executable cannot be empty")
	}
	command := strings.Split(*executable, " ")

	response, err := interactors.CreateJob(*serverURL, command)
	if err != nil {
		return err
	}

	_, _ = os.Stdout.WriteString(fmt.Sprintf("Created job successfully. Id: %s.\n", *response))
	return nil
}
