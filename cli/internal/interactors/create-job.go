package interactors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateJob(serverURL string, command []string) (*string, error) {
	postBody, err := json.Marshal(map[string][]string{
		"command": command,
	})
	if err != nil {
		return nil, err
	}

	httpResponse, err := http.Post(serverURL, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	defer closeBody(httpResponse.Body)

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	responseObject := string(body)
	return &responseObject, nil
}

func closeBody(response io.ReadCloser) {
	err := response.Close()
	if err != nil {
		fmt.Printf("failed to close http response body: %s\n", err)
	}
}
