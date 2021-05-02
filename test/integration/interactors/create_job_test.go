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
	"time"
)

type CreateJobInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *CreateJobInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
}

func (suite *CreateJobInteractorIntegrationTestSuite) SetupTest() {
	storage.CreateLogsDir()
}

func (suite *CreateJobInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldPersistJobWithCorrectParameters() {
	request := dto.CreateJobRequest{
		Command:          []string{"ls", "-la"},
		TimeoutInSeconds: 2,
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	job, err := repository.GetJob(response.ID)
	assert.Nil(suite.T(), err, "get job returned with error")

	assert.Equal(suite.T(), response.ID, job.ID, "persisted wrong ID")
	assert.Equal(suite.T(), []string{"ls", "-la"}, job.Command, "persisted wrong command")
	assert.Equal(suite.T(), time.Duration(2), job.TimeoutInSeconds, "persisted wrong timeout")
	assert.Equal(suite.T(), time.Time{}, job.FinishedAt, "persisted wrong finishedAt")
	assert.Equal(suite.T(), jobEntity.RUNNING, job.Status, "persisted wrong status")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldNotPersistJobWhenCreateProcessFails() {
	request := dto.CreateJobRequest{
		Command:          []string{},
		TimeoutInSeconds: 1,
	}

	response, err := interactors.CreateJob(request)
	assert.NotNil(suite.T(), err, "create job interactor returned without error")
	assert.Equal(suite.T(), response, dto.CreateJobResponse{}, "returned non empty job response")

	jobs, err := repository.GetAllJobs()
	assert.Nil(suite.T(), err, "get all jobs returned with error")
	assert.Equal(suite.T(), 0, len(jobs))
}

func TestCreateJobInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(CreateJobInteractorIntegrationTestSuite))
}