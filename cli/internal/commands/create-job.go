package commands

import (
	"cli/internal/config"
	"cli/internal/interactors"
	"errors"
	"flag"
	"fmt"
	"strings"
)

func CreateJob(parameters []string) error {
	execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := execCmd.String("s", config.GetDefaultServerURL(), "server url")
	executable := execCmd.String("c", "", "command to be executed")

	err := execCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *executable == "" {
		return errors.New("executable shouldn't be empty")
	}
	command := strings.Split(*executable, " ")

	response, err := interactors.CreateJob(*serverURL+"/jobs", command)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *response)
	return nil
}
