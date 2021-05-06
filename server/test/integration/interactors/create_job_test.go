package integration_interactors_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"fmt"
	"log"
	"os"

	"server/internal/dto"
	"server/internal/interactors"
	jobEntity "server/internal/models/job"
	"server/internal/repository"
	"server/internal/storage"
	"server/test/integration"
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
	err := storage.CreateLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test: %s", err))
	}
}

func (suite *CreateJobInteractorIntegrationTestSuite) TearDownTest() {
	err := integration.RollbackState()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to tear down test: %s", err))
	}
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldPersistJobWithCorrectParameters() {
	request := dto.CreateJobRequest{
		Command: []string{"ls", "-la"},
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	job, err := repository.GetJobOrFail(response.ID)
	assert.Nil(suite.T(), err, "get job returned with error")

	assert.Equal(suite.T(), response.ID, job.ID, "persisted wrong ID")
	assert.Equal(suite.T(), []string{"ls", "-la"}, job.Command, "persisted wrong command")
	assert.Nil(suite.T(), job.FinishedAt, "persisted wrong finishedAt")
	assert.Equal(suite.T(), jobEntity.RUNNING, job.Status, "persisted wrong status")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldNotPersistJobWhenCreateProcessFails() {
	request := dto.CreateJobRequest{
		Command: []string{},
	}

	response, err := interactors.CreateJob(request)
	assert.NotNil(suite.T(), err, "create job interactor returned without error")
	assert.Nil(suite.T(), response, "returned non empty job response")

	jobs, err := repository.GetAllJobs()
	assert.Nil(suite.T(), err, "get all jobs returned with error")
	assert.Equal(suite.T(), 0, len(jobs))
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldCreateOutputFilesSuccessfuly() {
	request := dto.CreateJobRequest{
		Command: []string{"ls", "-la"},
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	logFile, err := repository.GetLogFile(response.ID)
	assert.Nil(suite.T(), err, "get log file returned with error")
	assert.NotNil(suite.T(), logFile, "log file file is nil")
	defer closeFile(logFile)
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldNotCreateOutputFilesWhenCreateJobFails() {
	request := dto.CreateJobRequest{
		Command: []string{},
	}

	response, err := interactors.CreateJob(request)
	assert.NotNil(suite.T(), err, "create job interactor returned without error")
	assert.Nil(suite.T(), response, "returned non empty job response")

	filesNumber, err := integration.GetNumberOfLogFiles()
	assert.Nil(suite.T(), err, "get number of log files returned with error")
	assert.Equal(suite.T(), 0, *filesNumber, "number of log files should be 0")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestStdoutShouldHaveContentWhenProcessIsSuccessful() {
	request := dto.CreateJobRequest{
		Command: []string{"ls", "-la"},
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	logFile, err := repository.GetLogFile(response.ID)
	assert.Nil(suite.T(), err, "get log file returned with error")
	assert.NotNil(suite.T(), logFile, "log file is nil")
	defer closeFile(logFile)

	logFileInfo, err := logFile.Stat()
	assert.Nil(suite.T(), err, "get file info failed for log file")
	assert.Greater(suite.T(), logFileInfo.Size(), int64(0), "log file should have content")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestStderrShouldHaveContentWhenProcessFails() {
	request := dto.CreateJobRequest{
		Command: []string{"ls", "1000assa"},
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(250 * time.Millisecond)

	logFile, err := repository.GetLogFile(response.ID)
	assert.Nil(suite.T(), err, "get stderr file returned with error")
	assert.NotNil(suite.T(), logFile, "stderr file is nil")
	defer closeFile(logFile)

	logFileInfo, err := logFile.Stat()
	assert.Nil(suite.T(), err, "get file info failed for log file")
	assert.Greater(suite.T(), logFileInfo.Size(), int64(0), "log file should have content")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldPersistCorrectJobWhenProcessFailsExecution() {
	request := dto.CreateJobRequest{
		Command: []string{"ls", "100000asdas"},
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

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldPersistCorrectJobWhenProcessStateRemainsRunning() {
	request := dto.CreateJobRequest{
		Command: []string{"sleep", "4"},
	}

	response, err := interactors.CreateJob(request)
	assert.Nil(suite.T(), err, "create job interactor returned with error")

	time.Sleep(1100 * time.Millisecond)

	job, err := repository.GetJobOrFail(response.ID)
	assert.Nil(suite.T(), err, "get job returned with error")

	assert.Nil(suite.T(), job.FinishedAt, "persisted wrong finishedAt")
	assert.Equal(suite.T(), jobEntity.RUNNING, job.Status, "persisted wrong status")
	assert.Equal(suite.T(), -1, job.ExitCode, "persisted wrong exit code")
}

func (suite *CreateJobInteractorIntegrationTestSuite) TestShouldPersistCorrectJobWhenProcessSucceedsExecution() {
	request := dto.CreateJobRequest{
		Command: []string{"echo", "hello test world"},
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
