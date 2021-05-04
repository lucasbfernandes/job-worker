package worker

import (
	"errors"
	"io"
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

func NewProcess(command []string, timeoutInSeconds time.Duration) (*Process, error) {
	if len(command) == 0 || timeoutInSeconds <= 0 {
		return nil, errors.New("non empty command array and timeout greater than zero required")
	}
	execCmd := exec.Command(command[0], command[1:]...)

	return &Process{
		ExitChannel:      make(chan ExitReason, 1),
		TimeoutInSeconds: timeoutInSeconds,
		Command:          command,
		execCmd:          execCmd,
	}, nil
}

func (p *Process) SetStdoutWriter(stdoutWriter io.Writer) {
	p.execCmd.Stdout = stdoutWriter
}

func (p *Process) SetStderrWriter(stderrWriter io.Writer) {
	p.execCmd.Stderr = stderrWriter
}
