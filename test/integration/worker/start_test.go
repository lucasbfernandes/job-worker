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
	process, err := worker.NewProcess([]string{"ls", "-la"}, 2)
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	exitReason := <-process.ExitChannel
	assert.Equal(suite.T(), 0, exitReason.ExitCode, "Process should've returned with error code = 0")
}

func (suite *StartProcessIntegrationTestSuite) TestTimeoutExecutionShouldReturnCorrectStatus() {
	process, err := worker.NewProcess([]string{"sleep", "10"}, 2)
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	exitReason := <-process.ExitChannel
	assert.Equal(suite.T(), -1, exitReason.ExitCode, "Process should've returned with error code = -1")
}

func (suite *StartProcessIntegrationTestSuite) TestFailedExecutionShouldReturnCorrectStatus() {
	process, err := worker.NewProcess([]string{"ls", "10"}, 2)
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	exitReason := <-process.ExitChannel
	assert.Equal(suite.T(), 1, exitReason.ExitCode, "Process should've returned with error code = 1")
}

func TestStartProcessIntegrationTest(t *testing.T) {
	suite.Run(t, new(StartProcessIntegrationTestSuite))
}
