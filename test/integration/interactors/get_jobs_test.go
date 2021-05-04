package integration_interactors_test

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"job-worker/internal/dto"
	"job-worker/internal/interactors"
	jobEntity "job-worker/internal/models/job"
	"job-worker/internal/repository"
	"job-worker/internal/storage"
	"job-worker/test/integration"
	"testing"
)

type GetJobsInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *GetJobsInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
}

func (suite *GetJobsInteractorIntegrationTestSuite) SetupTest() {
	err := storage.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *GetJobsInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *GetJobsInteractorIntegrationTestSuite) TestShouldReturnEmptyArrayWhenNoJobsWerePersisted() {
	response, err := interactors.GetJobs()
	assert.Nil(suite.T(), err, "get jobs interactor returned with error")

	assert.Equal(suite.T(), &dto.GetJobsResponse{Jobs: []dto.JobResponse{}}, response, "wrong get jobs response")
	assert.Equal(suite.T(), 0, len(response.Jobs), "get jobs returned with wrong number of elements")
}

func (suite *GetJobsInteractorIntegrationTestSuite) TestShouldReturnCorrectArrayWhenOneJobIsPersisted() {
	job := jobEntity.NewJob([]string{"ls", "-la"}, 2)

	err := repository.UpsertJob(job)
	assert.Nil(suite.T(), err, "upsert job returned with error")

	response, err := interactors.GetJobs()
	assert.Nil(suite.T(), err, "get jobs interactor returned with error")

	expectedResponse := &dto.GetJobsResponse{Jobs: []dto.JobResponse{dto.JobResponseFromJob(job)}}
	assert.Equal(suite.T(), expectedResponse, response, "wrong get jobs response")
	assert.Equal(suite.T(), 1, len(response.Jobs), "get jobs returned with wrong number of elements")
}

func TestGetJobsInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobsInteractorIntegrationTestSuite))
}
