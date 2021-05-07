package interactors

import (
	"server/internal/dto"
	jobEntity "server/internal/models/job"
	userEntity "server/internal/models/user"
)

func (s *ServerInteractor) GetJobs(user *userEntity.User) (*dto.GetJobsResponse, error) {
	var jobs []*jobEntity.Job
	var err error

	if user.Role == userEntity.AdminRole {
		jobs, err = s.getAllJobs()
	} else {
		jobs, err = s.getUserJobs(user)
	}

	if err != nil {
		return nil, err
	}

	getJobsResponse := getJobResponses(jobs)
	return getJobsResponse, nil
}

func (s *ServerInteractor) getUserJobs(user *userEntity.User) ([]*jobEntity.Job, error) {
	jobs, err := s.Database.GetAllUserJobs(user.ID)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (s *ServerInteractor) getAllJobs() ([]*jobEntity.Job, error) {
	jobs, err := s.Database.GetAllJobs()
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func getJobResponses(jobs []*jobEntity.Job) *dto.GetJobsResponse {
	getJobsResponse := dto.GetJobsResponse{
		Jobs: make([]dto.JobResponse, 0),
	}
	for _, job := range jobs {
		getJobsResponse.Jobs = append(getJobsResponse.Jobs, dto.JobResponseFromJob(job))
	}
	return &getJobsResponse
}
