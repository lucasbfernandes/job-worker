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

// Exit code 124 was chosen to resemble the timeout command behaviour
// (i.e. https://man7.org/linux/man-pages/man1/timeout.1.html)
func (p *Process) handleTimeout() error {
	p.emitExitReason(124)
	err := p.execCmd.Process.Kill()
	if err != nil {
		log.Printf("failed to kill process after timeout: %s\n", err)
		return err
	}
	return nil
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
