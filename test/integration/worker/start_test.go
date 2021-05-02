package integration_worker_test

import (
	"github.com/stretchr/testify/assert"

	"job-worker/pkg/worker"

	"testing"
)

func TestSuccessfulExecutionShouldReturnCorrectStatus(t *testing.T) {
	process := worker.NewProcess([]string{"ls", "-la"}, 2)
	err := process.Start()
	if err != nil {
		t.Fatalf("Failed to start process")
	}
	exitReason := <-process.ExitChannel
	assert.Equal(t, 0, exitReason.ExitCode, "Process should've returned with error code = 0")
}

func TestTimeoutExecutionShouldReturnCorrectStatus(t *testing.T) {
	process := worker.NewProcess([]string{"sleep", "10"}, 2)
	err := process.Start()
	if err != nil {
		t.Fatalf("Failed to start process")
	}
	exitReason := <-process.ExitChannel
	assert.Equal(t, -1, exitReason.ExitCode, "Process should've returned with error code = -1")
}

func TestFailedExecutionShouldReturnCorrectStatus(t *testing.T) {
	process := worker.NewProcess([]string{"ls", "10"}, 2)
	err := process.Start()
	if err != nil {
		t.Fatalf("Failed to start process")
	}
	exitReason := <-process.ExitChannel
	assert.Equal(t, 1, exitReason.ExitCode, "Process should've returned with error code = 1")
}
