package integration_interactors_test

import (
	"fmt"
	"job-worker/internal/dto"
	jobEntity "job-worker/internal/models/job"
	"job-worker/internal/repository"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"job-worker/internal/interactors"
	"job-worker/internal/storage"
	"job-worker/test/integration"
	"testing"
)

type StopJobInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *StopJobInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
}

func (suite *StopJobInteractorIntegrationTestSuite) SetupTest() {
	err := storage.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *StopJobInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *StopJobInteractorIntegrationTestSuite) TestShouldReturnErrorWhenJobDoesNotExist() {
	err := interactors.StopJob("1233")
	assert.NotNil(suite.T(), err, "stop job interactor should return with error")
}

func (suite *StopJobInteractorIntegrationTestSuite) TestShouldStopJobSuccessfully() {
	request := dto.CreateJobRequest{
		Command: []string{"sleep", "10"},
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	err = interactors.StopJob(createJobResponse.ID)
	assert.Nil(suite.T(), err, "stop interactor should not return with error")

	time.Sleep(250 * time.Millisecond)

	job, err := repository.GetJobOrFail(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job should not return with error")

	assert.Equal(suite.T(), jobEntity.STOPPED, job.Status, "job status should be STOPPED")
	assert.Equal(suite.T(), -1, job.ExitCode, "job exit code should be -1")
}

func TestStopJobInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(StopJobInteractorIntegrationTestSuite))
}
