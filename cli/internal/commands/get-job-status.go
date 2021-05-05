package commands

import (
	"cli/internal/interactors"
	"errors"
	"flag"
	"fmt"
)

func GetJobStatus(parameters []string) error {
	statusCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := statusCmd.String("s", "", "server url")
	username := statusCmd.String("u", "", "username")
	jobID := statusCmd.String("i", "", "job id")

	err := statusCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *serverURL == "" || *username == "" || *jobID == "" {
		return errors.New("serverUrl, username and jobId shouldn't be empty")
	}

	response, err := interactors.GetJobStatus(*serverURL, *username, *jobID)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", *response)
	return nil
}
