package main

import (
	"cli/internal/commands"
	"os"
)

func main() {
	workerCLI := commands.NewWorkerCLI()
	err := workerCLI.ExecuteCommand(os.Args)
	if err != nil {
		_, _ = os.Stdout.WriteString("failed to execute command: " + err.Error() + "\n")
		os.Exit(1)
	}
	os.Exit(0)
}
