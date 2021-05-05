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

func (suite *StopProcessIntegrationTestSuite) TestShouldFailOnlyOneStartWhenConcurrentCallHappens() {
	process, err := worker.NewProcess([]string{"sleep", "5"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	var err1 error
	var err2 error
	firstStartChan := make(chan struct{}, 1)
	secondStartChan := make(chan struct{}, 1)

	go func() {
		err1 = process.Start()
		firstStartChan <- struct{}{}
	}()
	go func() {
		err2 = process.Start()
		secondStartChan <- struct{}{}
	}()

	<-firstStartChan
	<-secondStartChan
	close(firstStartChan)
	close(secondStartChan)

	assert.NotEqual(suite.T(), err1, err2, "errors must not be equal - only one should fail.")
}

func TestStartProcessIntegrationTest(t *testing.T) {
	suite.Run(t, new(StartProcessIntegrationTestSuite))
}
