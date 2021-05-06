package interactors

import (
	"server/internal/dto"
	"server/internal/repository"
)

func GetJobStatus(jobID string) (*dto.GetJobStatusResponse, error) {
	job, err := repository.GetJobOrFail(jobID)
	if err != nil {
		return nil, err
	}
	statusResponse := dto.JobStatusResponseFromJob(job)
	return &statusResponse, nil
}
