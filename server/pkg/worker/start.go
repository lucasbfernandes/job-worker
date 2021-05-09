package worker

import (
	"log"
	"time"
)

func (p *Process) Start() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	err := p.execCmd.Start()
	if err != nil {
		log.Printf("failed to start process: %s\n", err)
		return err
	}

	go p.waitExecution()

	return nil
}

func (p *Process) waitExecution() {
	err := p.execCmd.Wait()
	if err != nil {
		log.Printf("process finished with error: %s\n", err)
	}

	p.ExitChannel <- ExitReason{
		ExitCode:  p.execCmd.ProcessState.ExitCode(),
		Timestamp: time.Now(),
	}

	close(p.finishedChannel)
	close(p.ExitChannel)
}
