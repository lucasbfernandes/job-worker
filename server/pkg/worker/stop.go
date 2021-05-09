package worker

import (
	"errors"
	"fmt"
	"log"
	"time"
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

	select {
	case <-time.After(2 * time.Second):
		log.Printf("process stop channel timed out\n")
	case <-p.finishedChannel:
		log.Printf("process stopped successfully\n")
	}

	return nil
}
