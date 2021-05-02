package integration_worker_test

import (
	"github.com/stretchr/testify/assert"
	"job-worker/pkg/worker"
	"testing"
	"time"
)

func TestSuccessfulStopShouldReturnCorrectStatus(t *testing.T) {
	process, err := worker.NewProcess([]string{"sleep", "10"}, 20)
	assert.Nil(t, err, "Failed to create process.")

	err = process.Start()
	assert.Nil(t, err, "Process failed to start.")

	err = process.Stop()
	assert.Nil(t, err, "Process failed to stop.")

	exitReason := <-process.ExitChannel
	assert.Equal(t, -1, exitReason.ExitCode, "Process should've returned with error code = 1")
}

func TestShouldFailStopWhenProcessHasAlreadyStopped(t *testing.T) {
	process, err := worker.NewProcess([]string{"sleep", "5"}, 1)
	assert.Nil(t, err, "Failed to create process.")

	err = process.Start()
	assert.Nil(t, err, "Process failed to start.")

	time.Sleep(2 * time.Second)

	err = process.Stop()
	assert.NotNil(t, err, "Should have failed because process has already stopped.")
}

func TestShouldReturnErrorWhenProcessHasntStarted(t *testing.T) {
	process, err := worker.NewProcess([]string{"ls", "-la"}, 2)
	assert.Nil(t, err, "Failed to create process.")

	err = process.Stop()
	assert.NotNil(t, err, "Should have failed because process hasn't started yet.")
}
