package integration_interactors_test

import (
	"fmt"
	"server/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"server/internal/dto"
	"server/internal/interactors"
	"server/test/integration"
	"testing"
	"time"
)

type GetJobLogsInteractorIntegrationTestSuite struct {
	suite.Suite

	interactor *interactors.ServerInteractor

	adminToken string

	userToken string
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	suite.interactor, err = interactors.NewServerInteractor()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	suite.adminToken = "qTMaYIfw8q3esZ6Dv2rQ"
	suite.userToken = "9EzGJOTcMHFMXphfvAuM"
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) SetupTest() {
	err := repository.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}

	err = suite.interactor.Database.SeedUsers()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState(suite.interactor.Database)
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStderrIsEmpty() {
	request := dto.CreateJobRequest{
		Command: []string{"echo", "this is a test"},
	}

	createJobResponse, err := suite.interactor.CreateJob(request, suite.adminToken)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	getJobLogsResponse, err := suite.interactor.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "this is a test\n"
	assert.Equal(suite.T(), expectedLogs, *getJobLogsResponse, "")
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStdoutIsEmpty() {
	request := dto.CreateJobRequest{
		Command: []string{"ls", "abobora"},
	}

	createJobResponse, err := suite.interactor.CreateJob(request, suite.adminToken)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	getJobLogsResponse, err := suite.interactor.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "ls: abobora: No such file or directory\n"
	assert.Equal(suite.T(), expectedLogs, *getJobLogsResponse, "")
}

func (suite *GetJobLogsInteractorIntegrationTestSuite) TestShouldReturnLogsCorrectlyWhenStdoutAndStderrArentEmpty() {
	request := dto.CreateJobRequest{
		Command: []string{"sh", "-c", "echo hello test! && ls what"},
	}

	createJobResponse, err := suite.interactor.CreateJob(request, suite.adminToken)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	getJobLogsResponse, err := suite.interactor.GetJobLogs(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job logs interactor should not return with error")

	expectedLogs := "hello test!\nls: what: No such file or directory\n"
	assert.Equal(suite.T(), expectedLogs, *getJobLogsResponse, "")
}

func TestGetJobLogsInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobLogsInteractorIntegrationTestSuite))
}
