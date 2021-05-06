package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"cli/internal/config"
	"cli/internal/interactors"
	"testing"
)

type StopJobInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *StopJobInteractorIntegrationTestSuite) TestShouldNotReturnErrorWhenRequestIsSuccessful() {
	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	err := workerCLIInteractor.StopJob(config.GetDefaultServerURL(), "mock-id")
	assert.Nil(suite.T(), err, "error should be nil")
}

func TestStopJobInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(StopJobInteractorIntegrationTestSuite))
}
