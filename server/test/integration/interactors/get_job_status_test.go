package integration_interactors_test

import (
	"fmt"
	"server/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"server/internal/dto"
	"server/internal/interactors"
	"server/internal/models/job"
	"server/test/integration"
	"testing"
	"time"
)

type GetJobStatusInteractorIntegrationTestSuite struct {
	suite.Suite

	interactor *interactors.ServerInteractor
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	suite.interactor, err = interactors.NewServerInteractor()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) SetupTest() {
	err := repository.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState(suite.interactor.Database)
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnErrorWhenJobDoesNotExist() {
	response, err := suite.interactor.GetJobStatus("1233")
	assert.NotNil(suite.T(), err, "get job status interactor should return with error")
	assert.Nil(suite.T(), response, "wrong response for get job status")
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnCorrectStatusWhenJobSuccessfullyFinishes() {
	request := dto.CreateJobRequest{
		Command: []string{"echo", "hello test world"},
	}

	createJobResponse, err := suite.interactor.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	statusResponse, err := suite.interactor.GetJobStatus(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job status interactor should not return with error")

	assert.Equal(suite.T(), job.COMPLETED, statusResponse.Status, "wrong status, should be COMPLETED")
	assert.Equal(suite.T(), 0, statusResponse.ExitCode, "wrong exit code, should be 0")
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnCorrectStatusWhenJobFinishesWithError() {
	request := dto.CreateJobRequest{
		Command: []string{"cat", "hello test world"},
	}

	createJobResponse, err := suite.interactor.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	statusResponse, err := suite.interactor.GetJobStatus(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job status interactor should not return with error")

	assert.Equal(suite.T(), job.FAILED, statusResponse.Status, "wrong status, should be FAILED")
	assert.Equal(suite.T(), 1, statusResponse.ExitCode, "wrong exit code, should be 1")
}

func (suite *GetJobStatusInteractorIntegrationTestSuite) TestShouldReturnCorrectStatusWhenJobRemainsRunning() {
	request := dto.CreateJobRequest{
		Command: []string{"sleep", "10"},
	}

	createJobResponse, err := suite.interactor.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(1100 * time.Millisecond)

	statusResponse, err := suite.interactor.GetJobStatus(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job status interactor should not return with error")

	assert.Equal(suite.T(), job.RUNNING, statusResponse.Status, "wrong status, should be RUNNING")
	assert.Equal(suite.T(), -1, statusResponse.ExitCode, "wrong exit code, should be -1")
}

func TestGetJobStatusInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobStatusInteractorIntegrationTestSuite))
}
