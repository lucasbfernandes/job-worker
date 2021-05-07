package interactors

import (
	"server/internal/dto"
)

func (s *ServerInteractor) GetJobs() (*dto.GetJobsResponse, error) {
	getJobsResponse := dto.GetJobsResponse{
		Jobs: make([]dto.JobResponse, 0),
	}

	jobs, err := s.Database.GetAllJobs()
	if err != nil {
		return nil, err
	}

	for _, job := range jobs {
		getJobsResponse.Jobs = append(getJobsResponse.Jobs, dto.JobResponseFromJob(job))
	}
	return &getJobsResponse, nil
}
