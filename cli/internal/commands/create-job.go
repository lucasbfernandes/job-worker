package commands

import (
	"cli/internal/interactors"
	"errors"
	"flag"
	"fmt"
	"strings"
)

func CreateJob(parameters []string) error {
	execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := execCmd.String("s", "", "server url")
	username := execCmd.String("u", "", "username")
	executable := execCmd.String("c", "", "command that will be executed on the server")

	err := execCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *serverURL == "" || *username == "" || *executable == "" {
		return errors.New("serverUrl, username and executable shouldn't be empty")
	}

	command := strings.Split(*executable, " ")
	response, err := interactors.CreateJob(*serverURL, *username, command)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", *response)
	return nil
}
