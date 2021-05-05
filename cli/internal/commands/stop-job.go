package commands

import (
	"cli/internal/config"
	"cli/internal/interactors"
	"errors"
	"flag"
)

func StopJob(parameters []string) error {
	stopCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := stopCmd.String("s", config.GetDefaultServerURL(), "server url")
	jobID := stopCmd.String("i", "", "job id")

	err := stopCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *serverURL == "" || *jobID == "" {
		return errors.New("serverUrl, username and jobId shouldn't be empty")
	}

	err = interactors.StopJob(*serverURL, *jobID)
	if err != nil {
		return err
	}

	return nil
}
