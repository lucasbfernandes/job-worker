package integration_worker_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"job-worker/pkg/worker"
	"testing"
	"time"
)

type StopProcessIntegrationTestSuite struct {
	suite.Suite
}

func (suite *StopProcessIntegrationTestSuite) TestSuccessfulStopShouldReturnCorrectStatus() {
	process, err := worker.NewProcess([]string{"sleep", "10"}, 20)
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	err = process.Stop()
	assert.Nil(suite.T(), err, "Process failed to stop.")

	exitReason := <-process.ExitChannel
	assert.Equal(suite.T(), -1, exitReason.ExitCode, "Process should've returned with error code = 1")
}

func (suite *StopProcessIntegrationTestSuite) TestShouldFailStopWhenProcessHasAlreadyStopped() {
	process, err := worker.NewProcess([]string{"sleep", "5"}, 1)
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	time.Sleep(1100 * time.Millisecond)

	err = process.Stop()
	assert.NotNil(suite.T(), err, "Should have failed because process has already stopped.")
}

func (suite *StopProcessIntegrationTestSuite) TestShouldReturnErrorWhenProcessHasntStarted() {
	process, err := worker.NewProcess([]string{"ls", "-la"}, 2)
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Stop()
	assert.NotNil(suite.T(), err, "Should have failed because process hasn't started yet.")
}

func TestStopProcessIntegrationTest(t *testing.T) {
	suite.Run(t, new(StopProcessIntegrationTestSuite))
}
