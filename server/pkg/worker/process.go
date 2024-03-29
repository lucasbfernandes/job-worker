package worker

import (
	"errors"
	"io"
	"os/exec"
	"sync"
	"time"
)

type ExitReason struct {
	ExitCode int

	Timestamp time.Time
}

type Process struct {
	ExitChannel chan ExitReason

	Command []string

	execCmd *exec.Cmd

	finishedChannel chan struct{}

	mutex sync.Mutex
}

func NewProcess(command []string) (*Process, error) {
	if len(command) == 0 {
		return nil, errors.New("command array cannot be empty")
	}
	execCmd := exec.Command(command[0], command[1:]...)

	return &Process{
		ExitChannel:     make(chan ExitReason, 1),
		Command:         command,
		execCmd:         execCmd,
		finishedChannel: make(chan struct{}),
	}, nil
}

func (p *Process) SetStdoutWriter(stdoutWriter io.Writer) {
	p.execCmd.Stdout = stdoutWriter
}

func (p *Process) SetStderrWriter(stderrWriter io.Writer) {
	p.execCmd.Stderr = stderrWriter
}
