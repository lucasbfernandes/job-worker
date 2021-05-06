package commands

import (
	"cli/internal/interactors"
	"errors"
	"flag"
	"fmt"
)

func (w *WorkerCLI) GetJobs(parameters []string) error {
	getCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := getCmd.String("s", "", "server url")

	err := getCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *serverURL == "" {
		return errors.New("serverUrl and username shouldn't be empty")
	}

	response, err := interactors.GetJobs(*serverURL)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *response)
	return nil
}
