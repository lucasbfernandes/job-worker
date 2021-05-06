package interactors

import (
	"fmt"
	"log"
	"server/internal/dto"
	jobEntity "server/internal/models/job"
	"server/internal/repository"
	"server/pkg/worker"
	"time"
)

func CreateJob(createJobRequest dto.CreateJobRequest) (*dto.CreateJobResponse, error) {
	job := createJobRequest.ToJob()

	process, err := createWorkerProcess(createJobRequest.Command, job.ID)
	if err != nil {
		log.Printf("could not create process %s\n", err)
		return nil, err
	}

	savedJob, err := persistJob(job, process)
	if err != nil {
		log.Printf("could not persist job %s\n", err)
		return nil, err
	}

	err = process.Start()
	if err != nil {
		log.Printf("could not start process %s\n", err)
		setJobStatusFailed(savedJob)
		return nil, fmt.Errorf("couldn't start process for job %s with error: %s", savedJob.ID, err)
	}

	setJobStatusRunning(savedJob)
	go waitForExitReason(savedJob, process)

	return &dto.CreateJobResponse{ID: savedJob.ID}, nil
}

func waitForExitReason(job *jobEntity.Job, process *worker.Process) {
	exitReason := <-process.ExitChannel

	switch exitReason.ExitCode {
	case -1:
		finishJobWithStatusAndCode(*job, jobEntity.STOPPED, exitReason.ExitCode)
	case 0:
		finishJobWithStatusAndCode(*job, jobEntity.COMPLETED, exitReason.ExitCode)
	default:
		finishJobWithStatusAndCode(*job, jobEntity.FAILED, exitReason.ExitCode)
	}

	close(process.ExitChannel)
}

func persistJob(job *jobEntity.Job, process *worker.Process) (*jobEntity.Job, error) {
	job.SetProcess(process)

	err := repository.UpsertJob(job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func createWorkerProcess(command []string, jobID string) (*worker.Process, error) {
	process, err := worker.NewProcess(command)
	if err != nil {
		return nil, err
	}
	err = createWorkerProcessOutputFiles(process, jobID)
	if err != nil {
		return nil, err
	}
	return process, nil
}

func createWorkerProcessOutputFiles(process *worker.Process, jobID string) error {
	outfile, err := repository.CreateLogFile(jobID)
	if err != nil {
		return err
	}
	process.SetStdoutWriter(outfile)
	process.SetStderrWriter(outfile)
	return nil
}

func finishJobWithStatusAndCode(job jobEntity.Job, status string, exitCode int) {
	finishedAt := time.Now()

	job.Status = status
	job.FinishedAt = &finishedAt
	job.ExitCode = exitCode

	err := repository.UpsertJob(&job)
	if err != nil {
		log.Printf("failed to update job with status %s: %s\n", status, err)
	}
}

func setJobStatusRunning(job *jobEntity.Job) {
	job.Status = jobEntity.RUNNING
	err := repository.UpsertJob(job)
	if err != nil {
		log.Printf("failed to update job status to running: %s\n", err)
	}
}

func setJobStatusFailed(job *jobEntity.Job) {
	job.Status = jobEntity.FAILED
	err := repository.UpsertJob(job)
	if err != nil {
		log.Printf("failed to update job status to failed: %s\n", err)
	}
}
