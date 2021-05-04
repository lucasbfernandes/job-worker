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
		p.handleFinishedExecution(err)
	}()
}

func (p *Process) handleFinishedExecution(err error) {
	p.emitExitReason(p.execCmd.ProcessState.ExitCode())
	if err != nil {
		log.Printf("process finished with error: %s\n", err)
	}
}

func (p *Process) emitExitReason(exitCode int) {
	p.ExitChannel <- ExitReason{
		ExitCode:  exitCode,
		Timestamp: time.Now(),
	}
}
