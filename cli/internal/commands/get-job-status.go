package commands

import (
	"cli/internal/config"
	"cli/internal/interactors"
	"errors"
	"flag"
	"fmt"
)

func GetJobStatus(parameters []string) error {
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

	response, err := interactors.GetJobStatus(*serverURL, *jobID)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *response)
	return nil
}
