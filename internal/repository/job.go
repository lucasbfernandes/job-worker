package repository

import (
	"job-worker/internal/database"
	"job-worker/internal/models/job"
	"log"
)

func CreateJob(job job.Job) error {
	db := database.GetDB()
	txn := db.Txn(true)
	err := txn.Insert("job", job)
	if err != nil {
		log.Printf("failed to insert job: %s\n", err)
		return err
	}
	txn.Commit()
	return nil
}
