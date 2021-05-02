package integration_worker_test

import (
	"github.com/stretchr/testify/assert"

	"job-worker/pkg/worker"
	"testing"
	"time"
)

func TestShouldCreateProcessWithCorrectParameters(t *testing.T) {
	process, err := worker.NewProcess([]string{"ls", "-la"}, 2)
	assert.Nil(t, err, "Failed to create process.")

	assert.NotNil(t, process.StdoutPipe, "StdoutPipe is nil.")
	assert.NotNil(t, process.StderrPipe, "StderrPipe is nil.")
	assert.NotNil(t, process.ExitChannel, "ExitChannel is nil.")
	assert.Equal(t, []string{"ls", "-la"}, process.Command, "Invalid value for Command.")
	assert.Equal(t, time.Duration(2), process.TimeoutInSeconds, "Invalid value for TimeoutInSeconds.")
}

func TestShouldFailCreateProcessWhenCommandIsEmpty(t *testing.T) {
	_, err := worker.NewProcess([]string{}, 2)
	assert.NotNil(t, err, "Should fail process creation when command array is empty.")
}

func TestShouldNotFailCreateProcessWhenCommandHasNoParameters(t *testing.T) {
	_, err := worker.NewProcess([]string{"ls"}, 2)
	assert.Nil(t, err, "Should not fail process creation when command array has only one element.")
}

func TestShouldFailCreateProcessWithNonNaturalTimeoutNumber(t *testing.T) {
	_, err := worker.NewProcess([]string{"ls"}, -1)
	assert.NotNil(t, err, "Should fail process creation when negative timeout is provided.")

	_, err = worker.NewProcess([]string{"ls"}, 0)
	assert.NotNil(t, err, "Should fail process creation when zeroed timeout is provided.")
}
