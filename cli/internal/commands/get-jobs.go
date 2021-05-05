package commands

import (
	"cli/internal/interactors"
	"errors"
	"flag"
	"fmt"
)

func GetJobs(parameters []string) error {
	getCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := getCmd.String("s", "", "server url")
	username := getCmd.String("u", "", "username")

	err := getCmd.Parse(parameters)
	if err != nil {
		return err
	}

	if *serverURL == "" || *username == "" {
		return errors.New("serverUrl and username shouldn't be empty")
	}

	response, err := interactors.GetJobs(*serverURL, *username)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", *response)
	return nil
}
