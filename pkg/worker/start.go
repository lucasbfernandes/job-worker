package worker

import (
	"log"
	"time"
)

func (p *Process) Start() error {
	err := p.execCmd.Start()
	if err != nil {
		log.Printf("failed to start process: %s\n", err)
		return err
	}
	p.waitExecution()
	return nil
}

func (p *Process) waitExecution() {
	go func() {
		err := p.execCmd.Wait()
		if err != nil {
			log.Printf("process finished with error: %s\n", err)
		}
		p.finishedChannel <- struct{}{}
		p.ExitChannel <- ExitReason{
			ExitCode:  p.execCmd.ProcessState.ExitCode(),
			Timestamp: time.Now(),
		}
	}()
}
