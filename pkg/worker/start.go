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

// TODO check for goroutines leaks. Maybe implement goroutine timeout?
func (p *Process) waitExecution() {
	doneChannel := make(chan error, 1)

	go func() {
		doneChannel <- p.execCmd.Wait()
	}()

	go func() {
		select {
		case <-time.After(p.TimeoutInSeconds * time.Second):
			p.handleTimeout()

		case err := <-doneChannel:
			p.handleFinishedExecution(err)
		}
	}()
}

func (p *Process) handleTimeout() {
	err := p.execCmd.Process.Kill()
	if err != nil {
		// TODO maybe retry? Otherwise the goroutine may run forever
		log.Printf("failed to kill process after timeout: %s\n", err)
	}
	p.ExitChannel <- ExitReason{
		ExitCode:  p.execCmd.ProcessState.ExitCode(),
		Timestamp: time.Now(),
	}
}

func (p *Process) handleFinishedExecution(err error) {
	if err != nil {
		log.Printf("process finished with error: %s\n", err)
	}
	p.ExitChannel <- ExitReason{
		ExitCode:  p.execCmd.ProcessState.ExitCode(),
		Timestamp: time.Now(),
	}
}
