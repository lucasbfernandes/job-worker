package interactors

func (s *ServerInteractor) StopJob(jobID string) error {
	job, err := s.Database.GetJobOrFail(jobID)
	if err != nil {
		return err
	}

	err = job.GetProcess().Stop()
	if err != nil {
		return err
	}
	return nil
}
