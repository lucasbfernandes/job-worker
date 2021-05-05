package worker

import (
	"errors"
	"log"
)

func (p *Process) Stop() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.execCmd.Process == nil {
		return errors.New("process hasn't started yet")
	}

	err := p.execCmd.Process.Kill()
	if err != nil {
		log.Printf("failed to kill process: %s\n", err)
		return err
	}

	<-p.finishedChannel
	close(p.finishedChannel)

	return nil
}
