package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"cli/internal/config"
	"cli/internal/interactors"
	"testing"
)

type GetJobLogsInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnCorrectStringWhenRequestIsSuccessful() {
	expectedResponse := `"hello test!
 ls: wrongfile: No such file or directory"`

	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	response, err := workerCLIInteractor.GetJobLogs(config.GetDefaultServerURL(), "mock-id")
	assert.Nil(suite.T(), err, "error should be nil")
	assert.Equal(suite.T(), expectedResponse, *response, "wrong get logs response")
}

func TestGetJobLogsInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobLogsInteractorIntegrationTestSuite))
}
