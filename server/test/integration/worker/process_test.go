package integration_worker_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"server/pkg/worker"
	"testing"
)

type CreateProcessIntegrationTestSuite struct {
	suite.Suite
}

func (suite *CreateProcessIntegrationTestSuite) TestShouldCreateProcessWithCorrectParameters() {
	process, err := worker.NewProcess([]string{"ls", "-la"})
	assert.Nil(suite.T(), err, "Failed to create process.")

	assert.NotNil(suite.T(), process.ExitChannel, "ExitChannel is nil.")
	assert.Equal(suite.T(), []string{"ls", "-la"}, process.Command, "Invalid value for Command.")
}

func (suite *CreateProcessIntegrationTestSuite) TestShouldFailCreateProcessWhenCommandIsEmpty() {
	_, err := worker.NewProcess([]string{})
	assert.NotNil(suite.T(), err, "Should fail process creation when command array is empty.")
}

func (suite *CreateProcessIntegrationTestSuite) TestShouldNotFailCreateProcessWhenCommandHasNoParameters() {
	_, err := worker.NewProcess([]string{"ls"})
	assert.Nil(suite.T(), err, "Should not fail process creation when command array has only one element.")
}

func TestCreateProcessIntegrationTest(t *testing.T) {
	suite.Run(t, new(CreateProcessIntegrationTestSuite))
}
