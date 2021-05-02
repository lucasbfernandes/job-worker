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
	doneChannel := make(chan error, 1)

	go func() {
		doneChannel <- p.execCmd.Wait()
	}()

	go func() {
		select {
		case <-time.After(p.TimeoutInSeconds * time.Second):
			err := p.handleTimeout()
			if err != nil {
				close(doneChannel)
			}

		case err := <-doneChannel:
			p.handleFinishedExecution(err)
		}
	}()
}

func (p *Process) handleTimeout() error {
	err := p.execCmd.Process.Kill()
	if err != nil {
		log.Printf("failed to kill process after timeout: %s\n", err)
		return err
	}
	p.emitExitReason()
	return nil
}

func (p *Process) handleFinishedExecution(err error) {
	if err != nil {
		log.Printf("process finished with error: %s\n", err)
	}
	p.emitExitReason()
}

func (p *Process) emitExitReason() {
	p.ExitChannel <- ExitReason{
		ExitCode:  p.execCmd.ProcessState.ExitCode(),
		Timestamp: time.Now(),
	}
}

