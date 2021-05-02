package repository

import (
	"job-worker/internal/database"
	jobEntity "job-worker/internal/models/job"
	"log"
)

func UpsertJob(job jobEntity.Job) error {
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

func DeleteAllJobs() error {
	db := database.GetDB()
	txn := db.Txn(true)
	_, err := txn.DeleteAll("job", "id")
	if err != nil {
		log.Printf("failed to delete all jobs: %s\n", err)
		return err
	}
	txn.Commit()
	return nil
}

func GetJob(id string) (jobEntity.Job, error) {
	db := database.GetDB()
	txn := db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("job", "id", id)
	if err != nil {
		log.Printf("failed to delete all jobs: %s\n", err)
		return jobEntity.Job{}, nil
	}

	return raw.(jobEntity.Job), nil
}

func GetAllJobs() ([]jobEntity.Job, error) {
	jobs := make([]jobEntity.Job, 0)

	db := database.GetDB()
	txn := db.Txn(false)
	defer txn.Abort()

	jobIterator, err := txn.Get("job", "id")
	if err != nil {
		log.Printf("failed to get all jobs: %s\n", err)
		return []jobEntity.Job{}, err
	}

	for rawJob := jobIterator.Next(); rawJob != nil; rawJob = jobIterator.Next() {
		job := rawJob.(jobEntity.Job)
		jobs = append(jobs, job)
	}

	return jobs, nil
}
