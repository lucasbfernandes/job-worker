package commands

import (
	"cli/internal/config"
	"flag"
	"fmt"
)

func (w *WorkerCLI) GetJobs(parameters []string) error {
	getCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := getCmd.String("s", config.GetDefaultServerURL(), "server url")

	err := getCmd.Parse(parameters)
	if err != nil {
		return err
	}

	formattedJobs, err := w.workerCLIInteractor.GetJobs(*serverURL)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", *formattedJobs)
	return nil
}
