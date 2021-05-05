package main

import (
	"cli/internal/commands"
	"fmt"
	"os"
)

func handleInvalidCommand() {
	fmt.Println("expected 'exec', 'list', 'stop', 'status' or 'logs' subcommands")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		handleInvalidCommand()
	}

	parameters := os.Args[2:]
	switch os.Args[1] {

	case "exec":
		commands.CreateJob(parameters)
	case "list":
		commands.GetJobs(parameters)
	case "stop":
		commands.StopJob(parameters)
	case "status":
		commands.GetJobStatus(parameters)
	case "logs":
		commands.GetJobLogs(parameters)
	default:
		handleInvalidCommand()
	}
}
