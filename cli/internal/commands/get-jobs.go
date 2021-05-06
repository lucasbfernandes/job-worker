package commands

import (
	"cli/internal/config"
	"flag"
	"os"
)

func (w *WorkerCLI) GetJobs(parameters []string) error {
	getCmd := flag.NewFlagSet("exec", flag.ExitOnError)
	serverURL := getCmd.String("s", config.GetDefaultServerURL(), "server url")

	err := getCmd.Parse(parameters)
	if err != nil {
		return err
	}

	response, err := w.workerCLIInteractor.GetJobs(*serverURL)
	if err != nil {
		return err
	}

	_, _ = os.Stdout.WriteString(*response + "\n")
	return nil
}
