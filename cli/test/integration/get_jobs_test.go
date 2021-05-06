package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"cli/internal/config"
	"cli/internal/interactors"
	"testing"
)

type GetJobsInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *GetJobsInteractorIntegrationTestSuite) TestShouldReturnCorrectStringWhenGetSucceeds() {
	expectedResponse := `
1
id: ad94eaae-b33e-42f8-927d-c13c0fc4a1f3
command: [sh -c echo hello world]
status: COMPLETED
createdAt: 2021-05-04 22:12:09.733285 -0300 -03
finishedAt: 2021-05-04 22:12:09.745211 -0300 -03

2
id: 4321cafb-0749-4a8e-99ca-03bb782a3381
command: [sh -c wrongcommand]
status: FAILED
createdAt: 2021-05-04 22:12:09.733285 -0300 -03
finishedAt: 2021-05-04 22:12:09.745211 -0300 -03
`

	workerCLIInteractor := interactors.NewWorkerCLIInteractor()
	response, err := workerCLIInteractor.GetJobs(config.GetDefaultServerURL())
	assert.Nil(suite.T(), err, "error should be nil")
	assert.Equal(suite.T(), expectedResponse, *response, "wrong get logs response")
}

func TestGetJobsInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobsInteractorIntegrationTestSuite))
}
