package integration_worker_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"job-worker/pkg/worker"
	"testing"
	"time"
)

type CreateProcessIntegrationTestSuite struct {
	suite.Suite
}

func (suite *CreateProcessIntegrationTestSuite) TestShouldCreateProcessWithCorrectParameters() {
	process, err := worker.NewProcess([]string{"ls", "-la"}, 2)
	assert.Nil(suite.T(), err, "Failed to create process.")

	assert.NotNil(suite.T(), process.ExitChannel, "ExitChannel is nil.")
	assert.Equal(suite.T(), []string{"ls", "-la"}, process.Command, "Invalid value for Command.")
	assert.Equal(suite.T(), time.Duration(2), process.TimeoutInSeconds, "Invalid value for TimeoutInSeconds.")
}

func (suite *CreateProcessIntegrationTestSuite) TestShouldFailCreateProcessWhenCommandIsEmpty() {
	_, err := worker.NewProcess([]string{}, 2)
	assert.NotNil(suite.T(), err, "Should fail process creation when command array is empty.")
}

func (suite *CreateProcessIntegrationTestSuite) TestShouldNotFailCreateProcessWhenCommandHasNoParameters() {
	_, err := worker.NewProcess([]string{"ls"}, 2)
	assert.Nil(suite.T(), err, "Should not fail process creation when command array has only one element.")
}

func (suite *CreateProcessIntegrationTestSuite) TestShouldFailCreateProcessWithNonNaturalTimeoutNumber() {
	_, err := worker.NewProcess([]string{"ls"}, -1)
	assert.NotNil(suite.T(), err, "Should fail process creation when negative timeout is provided.")

	_, err = worker.NewProcess([]string{"ls"}, 0)
	assert.NotNil(suite.T(), err, "Should fail process creation when zeroed timeout is provided.")
}

func TestCreateProcessIntegrationTest(t *testing.T) {
	suite.Run(t, new(CreateProcessIntegrationTestSuite))
}
