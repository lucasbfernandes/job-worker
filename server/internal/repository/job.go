package repository

import (
	"fmt"
	jobEntity "server/internal/models/job"
)

func (db *InMemoryDatabase) UpsertJob(job *jobEntity.Job) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	txn := db.instance.Txn(true)
	err := txn.Insert("job", job)
	if err != nil {
		return fmt.Errorf("failed to insert job: %s", err)
	}
	txn.Commit()
	return nil
}

func (db *InMemoryDatabase) DeleteAllJobs() error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	txn := db.instance.Txn(true)
	_, err := txn.DeleteAll("job", "id")
	if err != nil {
		return fmt.Errorf("failed to delete all jobs: %s", err)
	}
	txn.Commit()
	return nil
}

func (db *InMemoryDatabase) GetJobOrFail(id string) (*jobEntity.Job, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	txn := db.instance.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("job", "id", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %s", err)
	}
	if raw == nil {
		return nil, fmt.Errorf("could not find job with id %s", id)
	}

	return raw.(*jobEntity.Job), nil
}

func (db *InMemoryDatabase) GetAllJobs() ([]*jobEntity.Job, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	txn := db.instance.Txn(false)
	defer txn.Abort()

	jobs := make([]*jobEntity.Job, 0)
	jobIterator, err := txn.Get("job", "id")
	if err != nil {
		return nil, fmt.Errorf("failed to get all jobs: %s", err)
	}

	for rawJob := jobIterator.Next(); rawJob != nil; rawJob = jobIterator.Next() {
		job := rawJob.(*jobEntity.Job)
		jobs = append(jobs, job)
	}

	return jobs, nil
}
