package interactors

import (
	"bytes"
	"cli/internal/dto"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateJob(serverURL string, command []string) (*string, error) {
	createJobRequest := dto.NewCreateJobRequest(command)
	responseBody, err := requestCreateJob(serverURL, createJobRequest)
	if err != nil {
		return nil, err
	}

	var createJobResponse dto.CreateJobResponse
	err = json.Unmarshal(responseBody, &createJobResponse)
	if err != nil {
		return nil, err
	}

	return &createJobResponse.ID, nil
}

func requestCreateJob(serverURL string, createJobRequest *dto.CreateJobRequest) ([]byte, error) {
	postBody, err := json.Marshal(createJobRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}

	httpResponse, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
