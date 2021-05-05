package commands

import (
	"cli/internal/interactors"
	"errors"
	"flag"
	"fmt"
)

func GetJobLogs(parameters []string) error {
	logsCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := logsCmd.String("s", "", "server url")
	username := logsCmd.String("u", "", "username")
	jobID := logsCmd.String("i", "", "job id")

	err := logsCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *serverURL == "" || *username == "" || *jobID == "" {
		return errors.New("serverUrl, username and jobId shouldn't be empty")
	}

	response, err := interactors.GetJobLogs(*serverURL, *username, *jobID)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", *response)
	return nil
}
