package worker

import (
	"errors"
	"log"
)

func (p *Process) Stop() error {
	if p.execCmd.Process == nil {
		return errors.New("process hasn't started yet")
	}

	err := p.execCmd.Process.Kill()
	if err != nil {
		log.Printf("failed to kill process: %s\n", err)
		return err
	}
	return nil
}
