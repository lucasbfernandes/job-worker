package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"cli/internal/interactors"
	"cli/test/integration"
	"testing"
)

type GetJobLogsInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnCorrectStringWhenRequestIsSuccessful() {
	expectedResponse := `hello test! ls: wrongfile: No such file or directory`

	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	response, err := workerCLIInteractor.GetJobLogs(integration.GetDefaultTestsServerURL(), "mock-id", "qTMaYIfw8q3esZ6Dv2rQ")
	assert.Nil(suite.T(), err, "error should be nil")
	assert.Equal(suite.T(), expectedResponse, *response, "wrong get logs response")
}

func TestGetJobLogsInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobLogsInteractorIntegrationTestSuite))
}
