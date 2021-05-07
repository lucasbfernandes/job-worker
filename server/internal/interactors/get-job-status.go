package interactors

import (
	"server/internal/dto"
)

func (s *ServerInteractor) GetJobStatus(jobID string) (*dto.GetJobStatusResponse, error) {
	job, err := s.Database.GetJobOrFail(jobID)
	if err != nil {
		return nil, err
	}
	statusResponse := dto.JobStatusResponseFromJob(job)
	return &statusResponse, nil
}
