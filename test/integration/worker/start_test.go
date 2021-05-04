package integration_worker_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"job-worker/pkg/worker"
	"testing"
)

type StartProcessIntegrationTestSuite struct {
	suite.Suite
}

func (suite *StartProcessIntegrationTestSuite) TestSuccessfulExecutionShouldReturnCorrectStatus() {
	process, err := worker.NewProcess([]string{"ls", "-la"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	exitReason := <-process.ExitChannel
	assert.Equal(suite.T(), 0, exitReason.ExitCode, "Process should've returned with error code = 0")
}

func (suite *StartProcessIntegrationTestSuite) TestFailedExecutionShouldReturnCorrectStatus() {
	process, err := worker.NewProcess([]string{"ls", "10"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	exitReason := <-process.ExitChannel
	assert.Equal(suite.T(), 1, exitReason.ExitCode, "Process should've returned with error code = 1")
}

func (suite *StartProcessIntegrationTestSuite) TestShouldFailWithStartSequence() {
	process, err := worker.NewProcess([]string{"sleep", "5"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	err = process.Start()
	assert.NotNil(suite.T(), err, "Process should fail to start again.")
}

func TestStartProcessIntegrationTest(t *testing.T) {
	suite.Run(t, new(StartProcessIntegrationTestSuite))
}
