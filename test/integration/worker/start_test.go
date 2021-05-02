package integration_worker_test

import (
	"github.com/stretchr/testify/assert"

	"job-worker/pkg/worker"

	"testing"
)

func TestSuccessfulExecutionShouldReturnCorrectStatus(t *testing.T) {
	process, err := worker.NewProcess([]string{"ls", "-la"}, 2)
	assert.Nil(t, err, "Failed to create process.")

	err = process.Start()
	assert.Nil(t, err, "Process failed to start.")

	exitReason := <-process.ExitChannel
	assert.Equal(t, 0, exitReason.ExitCode, "Process should've returned with error code = 0")
}

func TestTimeoutExecutionShouldReturnCorrectStatus(t *testing.T) {
	process, err := worker.NewProcess([]string{"sleep", "10"}, 2)
	assert.Nil(t, err, "Failed to create process.")

	err = process.Start()
	assert.Nil(t, err, "Process failed to start.")

	exitReason := <-process.ExitChannel
	assert.Equal(t, -1, exitReason.ExitCode, "Process should've returned with error code = -1")
}

func TestFailedExecutionShouldReturnCorrectStatus(t *testing.T) {
	process, err := worker.NewProcess([]string{"ls", "10"}, 2)
	assert.Nil(t, err, "Failed to create process.")

	err = process.Start()
	assert.Nil(t, err, "Process failed to start.")

	exitReason := <-process.ExitChannel
	assert.Equal(t, 1, exitReason.ExitCode, "Process should've returned with error code = 1")
}
