package interactors

import (
	"log"
	"server/internal/repository"
)

func StopJob(jobID string) error {
	job, err := repository.GetJobOrFail(jobID)
	if err != nil {
		log.Printf("could not find job: %s\n", err)
		return err
	}

	err = job.GetProcess().Stop()
	if err != nil {
		log.Printf("could not stop process: %s\n", err)
		return err
	}
	return nil
}
