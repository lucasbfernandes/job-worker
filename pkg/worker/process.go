package worker

import (
	"errors"
	"io"
	"log"
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

	StdoutPipe io.ReadCloser

	StderrPipe io.ReadCloser

	execCmd *exec.Cmd
}

func NewProcess(command []string, timeoutInSeconds time.Duration) (Process, error) {
	if len(command) == 0 || timeoutInSeconds <= 0 {
		return Process{}, errors.New("non empty command array and timeout greater than zero required")
	}

	execCmd := exec.Command(command[0], command[1:]...)
	stderrPipe, err := execCmd.StderrPipe()
	if err != nil {
		log.Printf("failed to create stderr pipe")
		return Process{}, err
	}
	stdoutPipe, err := execCmd.StdoutPipe()
	if err != nil {
		log.Printf("failed to create stdout pipe")
		return Process{}, err
	}

	return Process{
		ExitChannel:      make(chan ExitReason, 1),
		TimeoutInSeconds: timeoutInSeconds,
		Command:          command,
		StdoutPipe:       stdoutPipe,
		StderrPipe:       stderrPipe,
		execCmd:          execCmd,
	}, nil
}
