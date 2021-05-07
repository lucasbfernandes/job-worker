package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"cli/internal/interactors"
	"cli/test/integration"
	"testing"
)

type GetJobStatusInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnCorrectStringWhenRequestIsSuccessful() {
	expectedResponse := `
status: FAILED
user: user1
createdAt: 2021-05-04 19:23:22.341
finishedAt: 2021-05-04 19:23:22.406
exitCode: 1
`
	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	response, err := workerCLIInteractor.GetJobStatus(integration.GetDefaultTestsServerURL(), "mock-id", "qTMaYIfw8q3esZ6Dv2rQ")
	assert.Nil(suite.T(), err, "error should be nil")
	assert.Equal(suite.T(), expectedResponse, *response, "wrong get status response")
}

func TestGetJobStatusInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobStatusInteractorIntegrationTestSuite))
}
