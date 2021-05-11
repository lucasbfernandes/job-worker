package integration_e2e_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/internal/controllers"
	userEntity "server/internal/models/user"
	"server/internal/repository"
	"server/test/integration"
	"testing"
)

type GetJobsE2EIntegrationTestSuite struct {
	suite.Suite

	tlsServer *httptest.Server

	admin *userEntity.User

	server *controllers.Server
}

func (suite *GetJobsE2EIntegrationTestSuite) SetupSuite() {
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

func (suite *GetJobsE2EIntegrationTestSuite) TearDownSuite() {
	err := repository.DeleteLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to teardown test suite: %s", err))
	}
	suite.tlsServer.Close()
}

func (suite *GetJobsE2EIntegrationTestSuite) TestShouldReturn401WhenAuthHeaderIsNotPresent() {
	resp, err := http.Get(
		fmt.Sprintf("%s/jobs", suite.tlsServer.URL),
	)
	if err == nil {
		_ = resp.Body.Close()
	}
	assert.Nil(suite.T(), err, "request error should be nil")
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.StatusCode, "request response should have 401 status")
}

func (suite *GetJobsE2EIntegrationTestSuite) TestShouldReturn401WhenAuthHeaderExistsButIsInvalid() {
	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/jobs", suite.tlsServer.URL),
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

func (suite *GetJobsE2EIntegrationTestSuite) TestShouldReturn404WhenAuthHeaderExistsButUserDoesnt() {
	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/jobs", suite.tlsServer.URL),
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

func (suite *GetJobsE2EIntegrationTestSuite) TestShouldReturn200WhenAuthHeaderExistsAndIsCorrect() {
	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/jobs", suite.tlsServer.URL),
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

func TestGetJobsE2EIntegrationTest(t *testing.T) {
	suite.Run(t, new(GetJobsE2EIntegrationTestSuite))
}
