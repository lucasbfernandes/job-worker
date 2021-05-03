package integration_interactors_test

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"job-worker/internal/dto"
	"job-worker/internal/interactors"
	"job-worker/internal/storage"
	"job-worker/test/integration"
	"testing"
	"time"
)

type GetJobLogsInteractorIntegrationTestSuite struct {
	suite.Suite
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) SetupTest() {
	storage.CreateLogsDir()
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStderrIsEmpty() {
	request := dto.CreateJobRequest{
		Command:          []string{"echo", "this is a test"},
		TimeoutInSeconds: 20,
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(2 * time.Second)

	getJobLogsResponse, err := interactors.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "stdout:\nthis is a test\n\nstderr:\n"
	assert.Equal(suite.T(), expectedLogs, getJobLogsResponse, "")
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStdoutIsEmpty() {
	request := dto.CreateJobRequest{
		Command:          []string{"ls", "abobora"},
		TimeoutInSeconds: 20,
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(2 * time.Second)

	getJobLogsResponse, err := interactors.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "stdout:\n\nstderr:\nls: abobora: No such file or directory\n"
	assert.Equal(suite.T(), expectedLogs, getJobLogsResponse, "")
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStdoutAndStderrArentEmpty() {
	request := dto.CreateJobRequest{
		Command:          []string{"sh", "-c", "echo hello test! && ls what"},
		TimeoutInSeconds: 20,
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(2 * time.Second)

	getJobLogsResponse, err := interactors.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "stdout:\nhello test!\n\nstderr:\nls: what: No such file or directory\n"
	assert.Equal(suite.T(), expectedLogs, getJobLogsResponse, "")
}

func TestGetJobLogsInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobLogsInteractorIntegrationTestSuite))
}
