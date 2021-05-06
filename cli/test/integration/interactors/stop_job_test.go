package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"cli/internal/interactors"
	"cli/test/integration"
	"testing"
)

type StopJobInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *StopJobInteractorIntegrationTestSuite) TestShouldNotReturnErrorWhenRequestIsSuccessful() {
	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	err := workerCLIInteractor.StopJob(integration.GetDefaultTestsServerURL(), "mock-id")
	assert.Nil(suite.T(), err, "error should be nil")
}

func TestStopJobInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(StopJobInteractorIntegrationTestSuite))
}
