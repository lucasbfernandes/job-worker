package commands

import (
	"cli/internal/interactors"
	"errors"
)

type WorkerCLI struct {
	workerCLIInteractor *interactors.WorkerCLIInteractor
}

func NewWorkerCLI() *WorkerCLI {
	return &WorkerCLI{
		workerCLIInteractor: interactors.NewWorkerCLIInteractor(),
	}
}

func (w *WorkerCLI) ExecuteCommand(args []string) error {
	if len(args) < 2 {
		return errors.New("expected one of 'exec', 'list', 'stop', 'status' or 'logs' commands")
	}

	parameters := args[2:]
	switch args[1] {

	case "exec":
		return w.CreateJob(parameters)
	case "list":
		return w.GetJobs(parameters)
	case "stop":
		return w.StopJob(parameters)
	case "status":
		return w.GetJobStatus(parameters)
	case "logs":
		return w.GetJobLogs(parameters)
	default:
		return errors.New("unknown command")
	}
}
