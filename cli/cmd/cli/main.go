package main

import (
	"cli/internal/commands"
	"os"
)

func main() {
	workerCLI := commands.NewWorkerCLI()
	err := workerCLI.ExecuteCommand(os.Args)
	if err != nil {
		_, _ = os.Stdout.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	os.Exit(0)
}
