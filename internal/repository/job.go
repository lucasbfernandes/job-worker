package repository

import (
	"fmt"
	jobEntity "job-worker/internal/models/job"
	"job-worker/internal/storage"
	"log"
)

func UpsertJob(job *jobEntity.Job) error {
	db := storage.GetDB()
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
	db := storage.GetDB()
	txn := db.Txn(true)
	_, err := txn.DeleteAll("job", "id")
	if err != nil {
		log.Printf("failed to delete all jobs: %s\n", err)
		return err
	}
	txn.Commit()
	return nil
}

func GetJobOrFail(id string) (*jobEntity.Job, error) {
	db := storage.GetDB()
	txn := db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("job", "id", id)
	if err != nil {
		log.Printf("failed to get job: %s\n", err)
		return nil, nil
	}
	if raw == nil {
		log.Printf("could not find job\n")
		return nil, fmt.Errorf("could not find job with id %s", id)
	}

	return raw.(*jobEntity.Job), nil
}

func GetAllJobs() ([]*jobEntity.Job, error) {
	jobs := make([]*jobEntity.Job, 0)

	db := storage.GetDB()
	txn := db.Txn(false)
	defer txn.Abort()

	jobIterator, err := txn.Get("job", "id")
	if err != nil {
		log.Printf("failed to get all jobs: %s\n", err)
		return nil, err
	}

	for rawJob := jobIterator.Next(); rawJob != nil; rawJob = jobIterator.Next() {
		job := rawJob.(*jobEntity.Job)
		jobs = append(jobs, job)
	}

	return jobs, nil
}
