package commands

import (
	"cli/internal/config"
	"cli/internal/interactors"
	"errors"
	"flag"
)

type WorkerCLI struct {
	workerCLIInteractor *interactors.WorkerCLIInteractor
}

func NewWorkerCLI() *WorkerCLI {
	return &WorkerCLI{
		workerCLIInteractor: interactors.NewWorkerCLIInteractor(),
	}
}

func (w *WorkerCLI) ParseAndExecuteCommand(args []string) error {
	if len(args) < 2 {
		return errors.New("expected one of 'exec', 'list', 'stop', 'status' or 'logs' commands")
	}
	parameters := args[2:]

	cmd := flag.NewFlagSet("worker-cli", flag.ExitOnError)
	serverURL := cmd.String("s", config.GetDefaultServerURL(), "server url")
	jobID := cmd.String("i", "", "job id")
	apiToken := cmd.String("t", "", "user api token")

	err := cmd.Parse(parameters)
	if err != nil {
		return err
	}

	switch args[1] {

	case "exec":
		return w.CreateJob(*serverURL, cmd.Args(), *apiToken)
	case "list":
		return w.GetJobs(*serverURL, *apiToken)
	case "stop":
		return w.StopJob(*serverURL, *jobID, *apiToken)
	case "status":
		return w.GetJobStatus(*serverURL, *jobID, *apiToken)
	case "logs":
		return w.GetJobLogs(*serverURL, *jobID, *apiToken)
	default:
		return errors.New("unknown command")
	}
}
