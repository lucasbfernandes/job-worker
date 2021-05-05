package main

import (
	"cli/internal/commands"
	"fmt"
	"os"
)

//Create Job: job-worker exec -s SERVER_URL -u USERNAME -c EXECUTABLE [ARG...]
//List Jobs: job-worker list -s SERVER_URL -u USERNAME
//Stop Job: job-worker stop -s SERVER_URL -u USERNAME -i JOB_ID
//Get Job Status: job-worker status -s SERVER_URL -u USERNAME -i JOB_ID
//Get Job Logs: job-worker logs -s SERVER_URL -u USERNAME -i JOB_ID

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
		err := commands.CreateJob(parameters)
		if err != nil {
			fmt.Printf("failed to create job: %s\n", err)
		}
	case "list":
		err := commands.GetJobs(parameters)
		if err != nil {
			fmt.Printf("failed to get jobs: %s\n", err)
		}
	case "stop":
		err := commands.StopJob(parameters)
		if err != nil {
			fmt.Printf("failed to stop job: %s\n", err)
		}
	case "status":
		err := commands.GetJobStatus(parameters)
		if err != nil {
			fmt.Printf("failed to get job status: %s\n", err)
		}
	case "logs":
		err := commands.GetJobLogs(parameters)
		if err != nil {
			fmt.Printf("failed to get job logs: %s\n", err)
		}
	default:
		handleInvalidCommand()
	}
}
