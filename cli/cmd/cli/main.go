package main

import (
	"cli/internal/commands"
	"fmt"
	"os"
)

func main() {
	workerCLI := commands.NewWorkerCLI()
	err := workerCLI.ExecuteCommand(os.Args)
	if err != nil {
		fmt.Printf("failed to execute command: %s\n", err.Error())
		os.Exit(1)
	}
}
