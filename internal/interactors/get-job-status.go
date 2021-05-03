package interactors

import (
	"job-worker/internal/dto"
	"job-worker/internal/repository"
	"log"
)

func GetJobStatus(jobID string) (dto.GetJobStatusResponse, error) {
	job, err := repository.GetJobOrFail(jobID)
	if err != nil {
		log.Printf("failed to get job status: %s\n", err)
		return dto.GetJobStatusResponse{}, err
	}
	return dto.JobStatusResponseFromJob(job), nil
}
