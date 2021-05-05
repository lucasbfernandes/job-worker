package commands

import (
	"cli/internal/config"
	"cli/internal/interactors"
	"flag"
	"fmt"
	"strings"
)

func CreateJob(parameters []string) {
	execCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := execCmd.String("s", config.GetDefaultServerURL(), "server url")
	executable := execCmd.String("c", "", "command to be executed")

	err := execCmd.Parse(parameters)
	if err != nil {
		fmt.Printf("failed to parse command line arguments\n")
		return
	}

	if *executable == "" {
		fmt.Printf("executable cannot be empty\n")
		return
	}
	command := strings.Split(*executable, " ")

	response, err := interactors.CreateJob(*serverURL+"/jobs", command)
	if err != nil {
		fmt.Printf("failed to create job with error: %s\n", err)
		return
	}

	fmt.Printf("created job with id: %s\n", *response)
}
