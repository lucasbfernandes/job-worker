package integration_interactors_test

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"job-worker/internal/dto"
	"job-worker/internal/interactors"
	"job-worker/internal/models/job"
	"job-worker/internal/storage"
	"job-worker/test/integration"
	"testing"
	"time"
)

type GetJobStatusInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) SetupTest() {
	err := storage.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnErrorWhenJobDoesNotExist() {
	response, err := interactors.GetJobStatus("1233")
	assert.NotNil(suite.T(), err, "get job status interactor should return with error")
	assert.Nil(suite.T(), response, "wrong response for get job status")
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnCorrectStatusWhenJobSuccessfullyFinishes() {
	request := dto.CreateJobRequest{
		Command: []string{"echo", "hello test world"},
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	statusResponse, err := interactors.GetJobStatus(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job status interactor should not return with error")

	assert.Equal(suite.T(), job.COMPLETED, statusResponse.Status, "wrong status, should be COMPLETED")
	assert.Equal(suite.T(), 0, statusResponse.ExitCode, "wrong exit code, should be 0")
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnCorrectStatusWhenJobFinishesWithError() {
	request := dto.CreateJobRequest{
		Command: []string{"cat", "hello test world"},
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	statusResponse, err := interactors.GetJobStatus(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job status interactor should not return with error")

	assert.Equal(suite.T(), job.FAILED, statusResponse.Status, "wrong status, should be FAILED")
	assert.Equal(suite.T(), 1, statusResponse.ExitCode, "wrong exit code, should be 1")
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnCorrectStatusWhenJobRemainsRunning() {
	request := dto.CreateJobRequest{
		Command: []string{"sleep", "10"},
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(1100 * time.Millisecond)

	statusResponse, err := interactors.GetJobStatus(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job status interactor should not return with error")

	assert.Equal(suite.T(), job.RUNNING, statusResponse.Status, "wrong status, should be RUNNING")
	assert.Equal(suite.T(), -1, statusResponse.ExitCode, "wrong exit code, should be -1")
}

func TestGetJobStatusInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobStatusInteractorIntegrationTestSuite))
}
