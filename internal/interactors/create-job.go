package interactors

import (
	"fmt"
	"job-worker/internal/dto"
	jobEntity "job-worker/internal/models/job"
	"job-worker/internal/repository"
	"job-worker/pkg/worker"
	"log"
	"os"
	"path"
	"time"
)

func CreateJob(createJobRequest dto.CreateJobRequest) (dto.CreateJobResponse, error) {
	job := createJobRequest.ToJob()

	process, err := createWorkerProcess(createJobRequest.Command, time.Duration(createJobRequest.TimeoutInSeconds), job.ID)
	if err != nil {
		log.Printf("could not create process %s\n", err)
		return dto.CreateJobResponse{}, err
	}

	savedJob, err := persistJob(job)
	if err != nil {
		log.Printf("could not persist job %s\n", err)
		return dto.CreateJobResponse{}, err
	}

	err = process.Start()
	if err != nil {
		log.Printf("could not start process %s\n", err)
		finishJobWithStatus(savedJob, jobEntity.FAILED)
		return dto.CreateJobResponse{}, fmt.Errorf("couldn't start process for job %s with error: %s", savedJob.ID, err)
	}

	// TODO trigger goroutine that will watch process exit reason channel

	setJobStatusRunning(savedJob)
	return dto.CreateJobResponse{ID: savedJob.ID}, nil
}

func persistJob(job jobEntity.Job) (jobEntity.Job, error) {
	err := repository.UpsertJob(job)
	if err != nil {
		return jobEntity.Job{}, err
	}
	return job, nil
}

func createWorkerProcess(command []string, timeoutInSeconds time.Duration, jobID string) (worker.Process, error) {
	process, err := worker.NewProcess(command, timeoutInSeconds)
	if err != nil {
		return worker.Process{}, err
	}
	err = createWorkerProcessOutputFiles(process, jobID)
	if err != nil {
		return worker.Process{}, err
	}
	return process, nil
}

func createWorkerProcessOutputFiles(process worker.Process, jobID string) error {
	logsDIR := os.Getenv("LOGS_DIR")

	stdout, err := os.Create(path.Join(logsDIR, jobID+"-stdout"))
	if err != nil {
		return err
	}
	stderr, err := os.Create(path.Join(logsDIR, jobID+"-stderr"))
	if err != nil {
		return err
	}

	process.SetStdoutWriter(stdout)
	process.SetStderrWriter(stderr)

	return nil
}

func finishJobWithStatus(job jobEntity.Job, status string) {
	job.Status = status
	job.FinishedAt = time.Now()
	err := repository.UpsertJob(job)
	if err != nil {
		log.Printf("failed to update job with status %s: %s\n", status, err)
	}
}

func setJobStatusRunning(job jobEntity.Job) {
	job.Status = jobEntity.RUNNING
	err := repository.UpsertJob(job)
	if err != nil {
		log.Printf("failed to update job status to running: %s\n", err)
	}
}
