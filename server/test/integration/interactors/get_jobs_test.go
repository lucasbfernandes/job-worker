package integration_interactors_test

import (
	"fmt"
	userEntity "server/internal/models/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"server/internal/dto"
	"server/internal/interactors"
	jobEntity "server/internal/models/job"
	"server/internal/repository"
	"server/test/integration"
	"testing"
)

type GetJobsInteractorIntegrationTestSuite struct {
	suite.Suite

	interactor *interactors.ServerInteractor
}

func (suite *GetJobsInteractorIntegrationTestSuite) SetupSuite() {
	err := integration.BootstrapTestEnvironment()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	suite.interactor, err = interactors.NewServerInteractor()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
}

func (suite *GetJobsInteractorIntegrationTestSuite) SetupTest() {
	err := repository.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *GetJobsInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState(suite.interactor.Database)
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *GetJobsInteractorIntegrationTestSuite) TestShouldReturnEmptyArrayWhenNoJobsWerePersisted() {
	response, err := suite.interactor.GetJobs()
	assert.Nil(suite.T(), err, "get jobs interactor returned with error")

	assert.Equal(suite.T(), &dto.GetJobsResponse{Jobs: []dto.JobResponse{}}, response, "wrong get jobs response")
	assert.Equal(suite.T(), 0, len(response.Jobs), "get jobs returned with wrong number of elements")
}

func (suite *GetJobsInteractorIntegrationTestSuite) TestShouldReturnCorrectArrayWhenOneJobIsPersisted() {
	user, err := userEntity.NewUser("test", "random-token", "ADMIN")
	assert.Nil(suite.T(), err, "new user returned with error")

	job := jobEntity.NewJob([]string{"ls", "-la"}, user.ID)

	err = suite.interactor.Database.UpsertJob(job)
	assert.Nil(suite.T(), err, "upsert job returned with error")

	response, err := suite.interactor.GetJobs()
	assert.Nil(suite.T(), err, "get jobs interactor returned with error")

	expectedResponse := &dto.GetJobsResponse{Jobs: []dto.JobResponse{dto.JobResponseFromJob(job)}}
	assert.Equal(suite.T(), expectedResponse, response, "wrong get jobs response")
	assert.Equal(suite.T(), 1, len(response.Jobs), "get jobs returned with wrong number of elements")
}

func TestGetJobsInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobsInteractorIntegrationTestSuite))
}
