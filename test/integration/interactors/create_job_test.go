package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"fmt"
	"log"
	"os"

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

	job, err := repository.GetJobOrFail(response.ID)
	assert.Nil(suite.T(), err, "get job returned with error")

	assert.Equal(suite.T(), response.ID, job.ID, "persisted wrong ID")
	assert.Equal(suite.T(), []string{"ls", "-la"}, job.Command, "persisted wrong command")
	assert.Equal(suite.T(), time.Duration(2), job.TimeoutInSeconds, "persisted wrong timeout")
	assert.Nil(suite.T(), job.FinishedAt, "persisted wrong finishedAt")
	assert.Equal(suite.T(), jobEntity.RUNNING, job.Status, "persisted wrong status")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldNotPersistJobWhenCreateProcessFails() {
	request := dto.CreateJobRequest{
		Command:          []string{},
		TimeoutInSeconds: 1,
	}

	response, err := interactors.CreateJob(request)
	assert.NotNil(suite.T(), err, "create job interactor returned without error")
	assert.Equal(suite.T(), dto.CreateJobResponse{}, response, "returned non empty job response")

	jobs, err := repository.GetAllJobs()
	assert.Nil(suite.T(), err, "get all jobs returned with error")
	assert.Equal(suite.T(), 0, len(jobs))
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldCreateOutputFilesSuccessfuly() {
	request := dto.CreateJobRequest{
		Command:          []string{"ls", "-la"},
		TimeoutInSeconds: 2,
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	stdoutFile, err := repository.GetStdoutLogFile(response.ID)
	assert.Nil(suite.T(), err, "get stdout file returned with error")
	assert.NotNil(suite.T(), stdoutFile, "stdout file is nil")
	defer closeFile(stdoutFile)

	stderrFile, err := repository.GetStderrLogFile(response.ID)
	assert.Nil(suite.T(), err, "get stderr file returned with error")
	assert.NotNil(suite.T(), stderrFile, "stderr file is nil")
	defer closeFile(stderrFile)
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldNotCreateOutputFilesWhenCreateJobFails() {
	request := dto.CreateJobRequest{
		Command:          []string{},
		TimeoutInSeconds: 1,
	}

	response, err := interactors.CreateJob(request)
	assert.NotNil(suite.T(), err, "create job interactor returned without error")
	assert.Equal(suite.T(), dto.CreateJobResponse{}, response, "returned non empty job response")

	stdoutFile, err := repository.GetStdoutLogFile(response.ID)
	assert.NotNil(suite.T(), err, "get stdout file returned without error")
	assert.Nil(suite.T(), stdoutFile, "stdout file is not nil")

	stderrFile, err := repository.GetStderrLogFile(response.ID)
	assert.NotNil(suite.T(), err, "get stderr file returned without error")
	assert.Nil(suite.T(), stderrFile, "stderr file is not nil")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestStdoutShouldHaveContentWhenProcessIsSuccessful() {
	request := dto.CreateJobRequest{
		Command:          []string{"ls", "-la"},
		TimeoutInSeconds: 1,
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	stdoutFile, err := repository.GetStdoutLogFile(response.ID)
	assert.Nil(suite.T(), err, "get stdout file returned with error")
	assert.NotNil(suite.T(), stdoutFile, "stdout file is nil")
	defer closeFile(stdoutFile)

	stdoutInfo, err := stdoutFile.Stat()
	assert.Nil(suite.T(), err, "get file info failed for stdout")
	assert.Greater(suite.T(), stdoutInfo.Size(), int64(0), "stdout should have content")

	stderrFile, err := repository.GetStderrLogFile(response.ID)
	assert.Nil(suite.T(), err, "get stderr file returned with error")
	assert.NotNil(suite.T(), stderrFile, "stderr file is nil")
	defer closeFile(stderrFile)

	stderrInfo, err := stderrFile.Stat()
	assert.Nil(suite.T(), err, "get file info failed for stderr")
	assert.Equal(suite.T(), int64(0), stderrInfo.Size(), "stderr shouldn't have content")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestStderrShouldHaveContentWhenProcessFails() {
	request := dto.CreateJobRequest{
		Command:          []string{"ls", "1000assa"},
		TimeoutInSeconds: 1,
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	stdoutFile, err := repository.GetStdoutLogFile(response.ID)
	assert.Nil(suite.T(), err, "get stdout file returned with error")
	assert.NotNil(suite.T(), stdoutFile, "stdout file is nil")
	defer closeFile(stdoutFile)

	stdoutInfo, err := stdoutFile.Stat()
	assert.Nil(suite.T(), err, "get file info failed for stdout")
	assert.Equal(suite.T(), int64(0), stdoutInfo.Size(), "stdout shouldn't have content")

	stderrFile, err := repository.GetStderrLogFile(response.ID)
	assert.Nil(suite.T(), err, "get stderr file returned with error")
	assert.NotNil(suite.T(), stderrFile, "stderr file is nil")
	defer closeFile(stderrFile)

	stderrInfo, err := stderrFile.Stat()
	assert.Nil(suite.T(), err, "get file info failed for stderr")
	assert.Greater(suite.T(), stderrInfo.Size(), int64(0), "stderr should have content")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldPersistCorrectJobWhenProcessFailsExecution() {
	request := dto.CreateJobRequest{
		Command:          []string{"ls", "100000asdas"},
		TimeoutInSeconds: 1,
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	job, err := repository.GetJobOrFail(response.ID)
	assert.Nil(suite.T(), err, "get job returned with error")

	assert.NotNil(suite.T(), job.FinishedAt, "persisted wrong finishedAt")
	assert.Equal(suite.T(), jobEntity.FAILED, job.Status, "persisted wrong status")
	assert.Equal(suite.T(), 1, job.ExitCode, "persisted wrong exit code")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldPersistCorrectJobWhenProcessTimeoutsExecution() {
	request := dto.CreateJobRequest{
		Command:          []string{"sleep", "4"},
		TimeoutInSeconds: 1,
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(1100 * time.Millisecond)

	job, err := repository.GetJobOrFail(response.ID)
	assert.Nil(suite.T(), err, "get job returned with error")

	assert.NotNil(suite.T(), job.FinishedAt, "persisted wrong finishedAt")
	assert.Equal(suite.T(), jobEntity.TIMEOUT, job.Status, "persisted wrong status")
	assert.Equal(suite.T(), 124, job.ExitCode, "persisted wrong exit code")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldPersistCorrectJobWhenProcessSucceedsExecution() {
	request := dto.CreateJobRequest{
		Command:          []string{"echo", "hello test world"},
		TimeoutInSeconds: 1,
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	job, err := repository.GetJobOrFail(response.ID)
	assert.Nil(suite.T(), err, "get job returned with error")

	assert.NotNil(suite.T(), job.FinishedAt, "persisted wrong finishedAt")
	assert.Equal(suite.T(), jobEntity.COMPLETED, job.Status, "persisted wrong status")
	assert.Equal(suite.T(), 0, job.ExitCode, "persisted wrong exit code")
}

func TestCreateJobInteractorIntegrationTest(t *testing.T) {
	suite.Run(t, new(CreateJobInteractorIntegrationTestSuite))
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Printf("failed to close file with error: %s\n", err)
	}
}
