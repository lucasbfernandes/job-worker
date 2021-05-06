package interactors

import (
	"server/internal/repository"
)

func StopJob(jobID string) error {
	job, err := repository.GetJobOrFail(jobID)
	if err != nil {
		return err
	}

	err = job.GetProcess().Stop()
	if err != nil {
		return err
	}
	return nil
}
