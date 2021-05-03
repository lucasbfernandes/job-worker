package interactors

import (
	"fmt"
	"job-worker/internal/dto"
	jobEntity "job-worker/internal/models/job"
	"job-worker/internal/repository"
	"job-worker/pkg/worker"
	"log"
	"time"
)

func CreateJob(createJobRequest dto.CreateJobRequest) (dto.CreateJobResponse, error) {
	job := createJobRequest.ToJob()

	process, err := createWorkerProcess(createJobRequest.Command, time.Duration(createJobRequest.TimeoutInSeconds), job.ID)
	if err != nil {
		log.Printf("could not create process %s\n", err)
		return dto.CreateJobResponse{}, err
	}

	savedJob, err := persistJob(job, process)
	if err != nil {
		log.Printf("could not persist job %s\n", err)
		return dto.CreateJobResponse{}, err
	}

	err = process.Start()
	if err != nil {
		log.Printf("could not start process %s\n", err)
		setJobStatusFailed(savedJob)
		return dto.CreateJobResponse{}, fmt.Errorf("couldn't start process for job %s with error: %s", savedJob.ID, err)
	}

	setJobStatusRunning(savedJob)
	go waitForExitReason(savedJob, process)

	return dto.CreateJobResponse{ID: savedJob.ID}, nil
}

// This will never wait forever because of timeout constraints inside the worker library.
func waitForExitReason(job jobEntity.Job, process worker.Process) {
	exitReason := <-process.ExitChannel

	switch exitReason.ExitCode {
	case -1:
		finishJobWithStatusAndCode(job, jobEntity.STOPPED, exitReason.ExitCode)
	case 0:
		finishJobWithStatusAndCode(job, jobEntity.COMPLETED, exitReason.ExitCode)
	case 1:
		finishJobWithStatusAndCode(job, jobEntity.FAILED, exitReason.ExitCode)
	case 124:
		finishJobWithStatusAndCode(job, jobEntity.TIMEOUT, exitReason.ExitCode)
	default:
		finishJobWithStatusAndCode(job, jobEntity.FAILED, exitReason.ExitCode)
	}
}

func persistJob(job jobEntity.Job, process worker.Process) (jobEntity.Job, error) {
	job.SetProcess(&process)

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
	stdout, err := repository.CreateStdoutLogFile(jobID)
	if err != nil {
		return err
	}
	stderr, err := repository.CreateStderrLogFile(jobID)
	if err != nil {
		return err
	}

	process.SetStdoutWriter(stdout)
	process.SetStderrWriter(stderr)

	return nil
}

func finishJobWithStatusAndCode(job jobEntity.Job, status string, exitCode int) {
	finishedAt := time.Now()

	job.Status = status
	job.FinishedAt = &finishedAt
	job.ExitCode = exitCode

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

func setJobStatusFailed(job jobEntity.Job) {
	job.Status = jobEntity.FAILED
	err := repository.UpsertJob(job)
	if err != nil {
		log.Printf("failed to update job status to failed: %s\n", err)
	}
}
