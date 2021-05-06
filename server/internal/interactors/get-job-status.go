package interactors

import (
	"log"
	"server/internal/dto"
	"server/internal/repository"
)

func GetJobStatus(jobID string) (*dto.GetJobStatusResponse, error) {
	job, err := repository.GetJobOrFail(jobID)
	if err != nil {
		log.Printf("failed to get job status: %s\n", err)
		return nil, err
	}
	statusResponse := dto.JobStatusResponseFromJob(job)
	return &statusResponse, nil
}
