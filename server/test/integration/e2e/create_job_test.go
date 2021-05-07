package integration_e2e_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/internal/controllers"
	"server/internal/dto"
	"server/internal/repository"
	"server/test/integration"
	"testing"
)

type CreateJobE2EIntegrationTestSuite struct {
	suite.Suite

	tlsServer *httptest.Server
}

func (suite *CreateJobE2EIntegrationTestSuite) SetupSuite() {
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

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	ginEngine := server.GetGinEngine()
	suite.tlsServer = httptest.NewTLSServer(ginEngine)
}

func (suite *CreateJobE2EIntegrationTestSuite) TearDownSuite() {
	err := repository.DeleteLogsDir()
	if err != nil {
		suite.FailNow(fmt.Sprintf("failed to teardown test suite: %s", err))
	}
	suite.tlsServer.Close()
}

func (suite *CreateJobE2EIntegrationTestSuite) TestShouldReturn401WhenAuthHeaderIsNotPresent() {
	requestObj := dto.CreateJobRequest{
		Command: []string{"ls"},
	}
	requestBytes, err := json.Marshal(requestObj)
	assert.Nil(suite.T(), err, "failed to marshal request obj")

	resp, err := http.Post(
		fmt.Sprintf("%s/jobs", suite.tlsServer.URL),
		"application/json",
		bytes.NewReader(requestBytes),
	)
	if err == nil {
		_ = resp.Body.Close()
	}
	assert.Nil(suite.T(), err, "request error should be nil")
	assert.Equal(suite.T(), http.StatusUnauthorized, resp.StatusCode, "request response should have 401 status")
}

func (suite *CreateJobE2EIntegrationTestSuite) TestShouldReturn401WhenAuthHeaderExistsButIsInvalid() {
	requestObj := dto.CreateJobRequest{
		Command: []string{"ls"},
	}
	requestBytes, err := json.Marshal(requestObj)
	assert.Nil(suite.T(), err, "failed to marshal request obj")

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/jobs", suite.tlsServer.URL),
		bytes.NewReader(requestBytes),
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

func (suite *CreateJobE2EIntegrationTestSuite) TestShouldReturn404WhenAuthHeaderExistsButUserDont() {
	requestObj := dto.CreateJobRequest{
		Command: []string{"ls"},
	}
	requestBytes, err := json.Marshal(requestObj)
	assert.Nil(suite.T(), err, "failed to marshal request obj")

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/jobs", suite.tlsServer.URL),
		bytes.NewReader(requestBytes),
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

func (suite *CreateJobE2EIntegrationTestSuite) TestShouldReturn201WhenAuthHeaderExistsAndIsCorrect() {
	requestObj := dto.CreateJobRequest{
		Command: []string{"ls"},
	}
	requestBytes, err := json.Marshal(requestObj)
	assert.Nil(suite.T(), err, "failed to marshal request obj")

	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/jobs", suite.tlsServer.URL),
		bytes.NewReader(requestBytes),
	)
	assert.Nil(suite.T(), err, "failed to create http request")
	req.Header.Set("Authorization", "Bearer "+integration.GetAdminUserMockAPIToken())

	resp, err := httpClient.Do(req)
	if err == nil {
		_ = resp.Body.Close()
	}
	assert.Nil(suite.T(), err, "request error should be nil")
	assert.Equal(suite.T(), http.StatusCreated, resp.StatusCode, "request response should have 201 status")
}

func TestCreateJobE2EIntegrationTest(t *testing.T) {
	suite.Run(t, new(CreateJobE2EIntegrationTestSuite))
}
