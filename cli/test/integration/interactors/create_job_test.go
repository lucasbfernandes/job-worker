package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"cli/internal/interactors"
	"cli/test/integration"
	"testing"
)

type CreateJobInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldReturnErrorWhenRequestFailed() {
	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	response, err := workerCLIInteractor.CreateJob(integration.GetDefaultTestsServerURL(), []string{}, "qTMaYIfw8q3esZ6Dv2rQ")
	assert.Nil(suite.T(), response, "response should be nil")
	assert.NotNil(suite.T(), err, "error should not be nil")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldReturnResponseWhenRequestSucceeds() {
	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	response, err := workerCLIInteractor.CreateJob(integration.GetDefaultTestsServerURL(), []string{"ls"}, "qTMaYIfw8q3esZ6Dv2rQ")
	assert.NotNil(suite.T(), response, "response should not be nil")
	assert.Nil(suite.T(), err, "error should be nil")
}

func TestCreateJobInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(CreateJobInteractorIntegrationTestSuite))
}
