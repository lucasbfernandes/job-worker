package worker

import (
	"os/exec"

	"time"
)

type ExitReason struct {
	ExitCode int

	Timestamp time.Time
}

type Process struct {
	ExitChannel chan ExitReason

	TimeoutInSeconds time.Duration

	Command []string

	execCmd *exec.Cmd
}

// TODO insert parameter validations
func NewProcess(command []string, timeoutInSeconds time.Duration) Process {
	return Process{
		ExitChannel:      make(chan ExitReason, 1),
		TimeoutInSeconds: timeoutInSeconds,
		Command:          command,
		execCmd:          exec.Command(command[0], command[1:]...),
	}
}
