package interactors

import (
	"log"
	"server/internal/dto"
	"server/internal/repository"
)

func GetJobs() (*dto.GetJobsResponse, error) {
	getJobsResponse := dto.GetJobsResponse{
		Jobs: make([]dto.JobResponse, 0),
	}

	jobs, err := repository.GetAllJobs()
	if err != nil {
		log.Printf("failed to get jobs: %s\n", err)
		return nil, err
	}

	for _, job := range jobs {
		getJobsResponse.Jobs = append(getJobsResponse.Jobs, dto.JobResponseFromJob(job))
	}
	return &getJobsResponse, nil
}