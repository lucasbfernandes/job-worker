package integration_worker_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"server/pkg/worker"
	"testing"
)

type StopProcessIntegrationTestSuite struct {
	suite.Suite
}

func (suite *StopProcessIntegrationTestSuite) TestSuccessfulStopShouldReturnCorrectStatus() {
	process, err := worker.NewProcess([]string{"sleep", "10"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	err = process.Stop()
	assert.Nil(suite.T(), err, "Process failed to stop.")

	exitReason := <-process.ExitChannel
	assert.Equal(suite.T(), -1, exitReason.ExitCode, "Process should've returned with error code = 1")
}

func (suite *StopProcessIntegrationTestSuite) TestShouldFailStopWhenProcessHasAlreadyStopped() {
	process, err := worker.NewProcess([]string{"sleep", "5"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	err = process.Stop()
	assert.Nil(suite.T(), err, "Process stop should not fail.")

	err = process.Stop()
	assert.NotNil(suite.T(), err, "Should have failed because process has already stopped.")
}

func (suite *StopProcessIntegrationTestSuite) TestShouldFailOnlyOneStopWhenConcurrentCallHappens() {
	process, err := worker.NewProcess([]string{"sleep", "5"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Start()
	assert.Nil(suite.T(), err, "Process failed to start.")

	var err1 error
	var err2 error
	firstStopChan := make(chan struct{}, 1)
	secondStopChan := make(chan struct{}, 1)

	go func() {
		err1 = process.Stop()
		firstStopChan <- struct{}{}
	}()
	go func() {
		err2 = process.Stop()
		secondStopChan <- struct{}{}
	}()

	<-firstStopChan
	<-secondStopChan
	close(firstStopChan)
	close(secondStopChan)

	assert.NotEqual(suite.T(), err1, err2, "errors must not be equal - only one should fail.")
}

func (suite *StopProcessIntegrationTestSuite) TestShouldReturnErrorWhenProcessHasntStarted() {
	process, err := worker.NewProcess([]string{"ls", "-la"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	err = process.Stop()
	assert.NotNil(suite.T(), err, "Should have failed because process hasn't started yet.")
}

func TestStopProcessIntegrationTest(t *testing.T) {
	suite.Run(t, new(StopProcessIntegrationTestSuite))
}
