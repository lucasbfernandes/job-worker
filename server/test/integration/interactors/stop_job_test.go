package integration_interactors_test

import (
	"fmt"
	"server/internal/dto"
	jobEntity "server/internal/models/job"
	userEntity "server/internal/models/user"
	"server/internal/repository"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"server/internal/interactors"
	"server/test/integration"
	"testing"
)

type StopJobInteractorIntegrationTestSuite struct {
	suite.Suite

	interactor *interactors.ServerInteractor

	admin *userEntity.User

	user *userEntity.User
}

func (suite *StopJobInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	suite.interactor, err = interactors.NewServerInteractor()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	err = suite.interactor.Database.SeedUsers()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	suite.admin, err = suite.interactor.Database.GetUserOrFailByAPIToken("qTMaYIfw8q3esZ6Dv2rQ")
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	suite.user, err = suite.interactor.Database.GetUserOrFailByAPIToken("9EzGJOTcMHFMXphfvAuM")
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
}

func (suite *StopJobInteractorIntegrationTestSuite) SetupTest() {
	err := repository.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *StopJobInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState(suite.interactor.Database)
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *StopJobInteractorIntegrationTestSuite) TestShouldReturnErrorWhenJobDoesNotExist() {
	err := suite.interactor.StopJob("1233")
	assert.NotNil(suite.T(), err, "stop job interactor should return with error")
}

func (suite *StopJobInteractorIntegrationTestSuite) TestShouldStopJobSuccessfully() {
	request := dto.CreateJobRequest{
		Command: []string{"sleep", "10"},
	}

	createJobResponse, err := suite.interactor.CreateJob(request, suite.admin)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	err = suite.interactor.StopJob(createJobResponse.ID)
	assert.Nil(suite.T(), err, "stop interactor should not return with error")

	time.Sleep(250 * time.Millisecond)

	job, err := suite.interactor.Database.GetJobOrFail(createJobResponse.ID)
	assert.Nil(suite.T(), err, "get job should not return with error")

	assert.Equal(suite.T(), jobEntity.STOPPED, job.Status, "job status should be STOPPED")
	assert.Equal(suite.T(), -1, job.ExitCode, "job exit code should be -1")
}

func (suite *StopJobInteractorIntegrationTestSuite) TestShouldFailWhenJobHasAlreadyFinished() {
	request := dto.CreateJobRequest{
		Command: []string{"sleep", "10"},
	}

	createJobResponse, err := suite.interactor.CreateJob(request, suite.admin)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	err = suite.interactor.StopJob(createJobResponse.ID)
	assert.Nil(suite.T(), err, "stop interactor should not return with error")

	err = suite.interactor.StopJob(createJobResponse.ID)
	assert.NotNil(suite.T(), err, "stop interactor should return with error")
}

func TestStopJobInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(StopJobInteractorIntegrationTestSuite))
}
