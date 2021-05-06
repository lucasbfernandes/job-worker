package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"cli/internal/config"
	"cli/internal/interactors"
	"testing"
)

type GetJobStatusInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnCorrectStringWhenRequestIsSuccessful() {
	expectedResponse := `
status: FAILED
createdAt: 2021-05-04 19:23:22.341245 -0300 -03
finishedAt: 2021-05-04 19:23:22.406849 -0300 -03
exitCode: 1
`
	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	response, err := workerCLIInteractor.GetJobStatus(config.GetDefaultServerURL(), "mock-id")
	assert.Nil(suite.T(), err, "error should be nil")
	assert.Equal(suite.T(), expectedResponse, *response, "wrong get status response")
}

func TestGetJobStatusInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobStatusInteractorIntegrationTestSuite))
}
