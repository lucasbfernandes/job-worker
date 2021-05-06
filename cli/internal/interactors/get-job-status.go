package interactors

import (
	"github.com/go-resty/resty/v2"

	"cli/internal/dto"
	"errors"
	"fmt"
)

func (i *WorkerCLIInteractor) GetJobStatus(serverURL string, jobID string) (*string, error) {
	getJobStatusResponse, err := requestGetJobStatus(serverURL, jobID)
	if err != nil {
		return nil, err
	}

	parsedResponse := parseGetJobStatusResponse(getJobStatusResponse)
	return parsedResponse, nil
}

func requestGetJobStatus(serverURL string, jobID string) (*dto.GetJobStatusResponse, error) {
	var getJobStatusResponse dto.GetJobStatusResponse
	var getJobStatusError dto.JobsError

	client := resty.New()
	response, err := client.R().
		SetResult(&getJobStatusResponse).
		SetError(&getJobStatusError).
		Get(serverURL + jobsPath + "/" + jobID + getJobStatusPath)

	if err != nil {
		return nil, err
	}

	if response.IsError() {
		if getJobStatusError.Error != "" {
			return nil, errors.New(getJobStatusError.Error)
		}
		return nil, errors.New("could not get job status")
	}

	return &getJobStatusResponse, nil
}

func parseGetJobStatusResponse(response *dto.GetJobStatusResponse) *string {
	parsedResponse := fmt.Sprintf(
		"\nstatus: %s\ncreatedAt: %s\nfinishedAt: %s\nexitCode: %d\n",
		response.Status, response.CreatedAt.Format(dateLayout), response.FinishedAt.Format(dateLayout), response.ExitCode,
	)
	return &parsedResponse
}
