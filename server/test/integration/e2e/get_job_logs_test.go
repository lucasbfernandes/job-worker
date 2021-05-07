package integration_e2e_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/internal/controllers"
	"server/internal/dto"
	jobEntity "server/internal/models/job"
	userEntity "server/internal/models/user"
	"server/internal/repository"
	"server/test/integration"
	"testing"
)

type GetJobLogsE2EIntegrationTestSuite struct {
	suite.Suite

	tlsServer *httptest.Server

	admin *userEntity.User

	server *controllers.Server
}

func (suite *GetJobLogsE2EIntegrationTestSuite) SetupSuite() {
	err := integration.SetTestJWTCertPath()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	server, err := controllers.NewServer()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	err = server.SetupState()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}
	suite.server = server

	suite.admin, err = server.Interactor.Database.GetUserOrFailByAPIToken("qTMaYIfw8q3esZ6Dv2rQ")
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to setup test suite: %s", err))
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	ginEngine := server.GetGinEngine()
	suite.tlsServer = httptest.NewTLSServer(ginEngine)
}

func (suite *GetJobLogsE2EIntegrationTestSuite) TearDownSuite() {
	err := repository.DeleteLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to teardown test suite: %s", err))
	}
	suite.tlsServer.Close()
}

func (suite *GetJobLogsE2EIntegrationTestSuite) TestShouldReturn401WhenAuthHeaderIsNotPresent() {
	job := jobEntity.NewJob([]string{"ls", "-la"}, suite.admin.ID)
	err := suite.server.Interactor.Database.UpsertJob(job)
	assert.Nil(suite.T(), err, "upsert job returned with error")

	resp, err := http.Get(
		fmt.Sprintf("%s/jobs/"+job.ID+"/logs", suite.tlsServer.URL),
	)
	if err == nil {
		_ = resp.Body.Close()
	}
	assert.Nil(suite.T(), err, "request error should be nil")
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.StatusCode, "request response should have 401 status")
}

func (suite *GetJobLogsE2EIntegrationTestSuite) TestShouldReturn401WhenAuthHeaderExistsButIsInvalid() {
	job := jobEntity.NewJob([]string{"ls", "-la"}, suite.admin.ID)
	err := suite.server.Interactor.Database.UpsertJob(job)
	assert.Nil(suite.T(), err, "upsert job returned with error")

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/jobs/"+job.ID+"/logs", suite.tlsServer.URL),
		nil,
	)
	assert.Nil(suite.T(), err, "failed to create http request")

	req.Header.Set("Authorization", "Bearer "+integration.GetInvalidBearerToken())

	resp, err := httpClient.Do(req)
	if err == nil {
		_ = resp.Body.Close()
	}

	assert.Nil(suite.T(), err, "request error should be nil")
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.StatusCode, "request response should have 401 status")
}

func (suite *GetJobLogsE2EIntegrationTestSuite) TestShouldReturn401WhenAuthHeaderExistsIsCorrectButFromDiffUser() {
	job := jobEntity.NewJob([]string{"ls", "-la"}, suite.admin.ID)
	err := suite.server.Interactor.Database.UpsertJob(job)
	assert.Nil(suite.T(), err, "upsert job returned with error")

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/jobs/"+job.ID+"/logs", suite.tlsServer.URL),
		nil,
	)
	assert.Nil(suite.T(), err, "failed to create http request")

	req.Header.Set("Authorization", "Bearer "+integration.GetUserMockAPIToken())

	resp, err := httpClient.Do(req)
	if err == nil {
		_ = resp.Body.Close()
	}

	assert.Nil(suite.T(), err, "request error should be nil")
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.StatusCode, "request response should have 401 status")
}

func (suite *GetJobLogsE2EIntegrationTestSuite) TestShouldReturn404WhenAuthHeaderExistsButUserDoesnt() {
	job := jobEntity.NewJob([]string{"ls", "-la"}, suite.admin.ID)
	err := suite.server.Interactor.Database.UpsertJob(job)
	assert.Nil(suite.T(), err, "upsert job returned with error")

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/jobs/"+job.ID+"/logs", suite.tlsServer.URL),
		nil,
	)
	assert.Nil(suite.T(), err, "failed to create http request")

	req.Header.Set("Authorization", "Bearer "+integration.GetNonExistentUserAPIToken())

	resp, err := httpClient.Do(req)
	if err == nil {
		_ = resp.Body.Close()
	}

	assert.Nil(suite.T(), err, "request error should be nil")
	assert.Equal(suite.T(), http.StatusNotFound, resp.StatusCode, "request response should have 404 status")
}

func (suite *GetJobLogsE2EIntegrationTestSuite) TestShouldReturn200WhenAuthHeaderExistsAndIsCorrect() {
	request := dto.CreateJobRequest{
		Command: []string{"ls", "-la"},
	}

	response, err := suite.server.Interactor.CreateJob(request, suite.admin)
	assert.Nil(suite.T(), err, "create job interactor returned with error")
	assert.NotNil(suite.T(), response, "create job interactor response cannot be nil")

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/jobs/"+response.ID+"/logs", suite.tlsServer.URL),
		nil,
	)
	assert.Nil(suite.T(), err, "failed to create http request")

	req.Header.Set("Authorization", "Bearer "+integration.GetAdminUserMockAPIToken())

	resp, err := httpClient.Do(req)
	if err == nil {
		_ = resp.Body.Close()
	}

	assert.Nil(suite.T(), err, "request error should be nil")
	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode, "request response should have 200 status")
}

func TestGetJobLogsE2EIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobLogsE2EIntegrationTestSuite))
}
