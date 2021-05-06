package integration_interactors_test

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"server/internal/dto"
	"server/internal/interactors"
	"server/internal/storage"
	"server/test/integration"
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
	err := storage.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStderrIsEmpty() {
	request := dto.CreateJobRequest{
		Command: []string{"echo", "this is a test"},
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	getJobLogsResponse, err := interactors.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "this is a test\n"
	assert.Equal(suite.T(), expectedLogs, *getJobLogsResponse, "")
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStdoutIsEmpty() {
	request := dto.CreateJobRequest{
		Command: []string{"ls", "abobora"},
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	getJobLogsResponse, err := interactors.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "ls: abobora: No such file or directory\n"
	assert.Equal(suite.T(), expectedLogs, *getJobLogsResponse, "")
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStdoutAndStderrArentEmpty() {
	request := dto.CreateJobRequest{
		Command: []string{"sh", "-c", "echo hello test! && ls what"},
	}

	createJobResponse, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	getJobLogsResponse, err := interactors.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "hello test!\nls: what: No such file or directory\n"
	assert.Equal(suite.T(), expectedLogs, *getJobLogsResponse, "")
}

func TestGetJobLogsInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobLogsInteractorIntegrationTestSuite))
}
