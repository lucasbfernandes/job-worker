package commands

import (
	"cli/internal/interactors"
	"errors"
	"flag"
)

func StopJob(parameters []string) error {
	stopCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := stopCmd.String("s", "", "server url")
	username := stopCmd.String("u", "", "username")
	jobID := stopCmd.String("i", "", "job id")

	err := stopCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *serverURL == "" || *username == "" || *jobID == "" {
		return errors.New("serverUrl, username and jobId shouldn't be empty")
	}

	err = interactors.StopJob(*serverURL, *username, *jobID)
	if err != nil {
		return err
	}
	return nil
}
