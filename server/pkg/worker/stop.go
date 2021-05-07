package worker

import (
	"errors"
	"fmt"
)

func (p *Process) Stop() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.execCmd.Process == nil {
		return errors.New("process hasn't started yet")
	}

	err := p.execCmd.Process.Kill()
	if err != nil {
		return fmt.Errorf("failed to kill process: %s", err)
	}

	<-p.finishedChannel

	return nil
}
