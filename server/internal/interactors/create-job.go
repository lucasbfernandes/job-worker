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

	process, err := createWorkerProcess(createJobRequest.Command)
	if err != nil {
		return nil, err
	}

	err = createWorkerProcessOutputFiles(process, job.ID)
	if err != nil {
		return nil, err
	}

	err = process.Start()
	if err != nil {
		return nil, err
	}

	savedJob, err := persistJob(job, process)
	if err != nil {
		return nil, err
	}

	err = setJobStatusRunning(savedJob)
	if err != nil {
		return nil, err
	}

	go waitForExitReason(savedJob, process)

	return &dto.CreateJobResponse{ID: savedJob.ID}, nil
}

func waitForExitReason(job *jobEntity.Job, process *worker.Process) {
	exitReason := <-process.ExitChannel
	var err error

	switch exitReason.ExitCode {
	case -1:
		err = finishJobWithStatusAndCode(*job, jobEntity.STOPPED, exitReason.ExitCode)
	case 0:
		err = finishJobWithStatusAndCode(*job, jobEntity.COMPLETED, exitReason.ExitCode)
	default:
		err = finishJobWithStatusAndCode(*job, jobEntity.FAILED, exitReason.ExitCode)
	}

	if err != nil {
		log.Printf("%s\n", err.Error())
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

func createWorkerProcess(command []string) (*worker.Process, error) {
	process, err := worker.NewProcess(command)
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

func finishJobWithStatusAndCode(job jobEntity.Job, status string, exitCode int) error {
	finishedAt := time.Now()

	job.Status = status
	job.FinishedAt = &finishedAt
	job.ExitCode = exitCode

	err := repository.UpsertJob(&job)
	if err != nil {
		return fmt.Errorf("failed to update job with status %s: %s", status, err)
	}
	return nil
}

func setJobStatusRunning(job *jobEntity.Job) error {
	job.Status = jobEntity.RUNNING
	err := repository.UpsertJob(job)
	if err != nil {
		return fmt.Errorf("failed to update job status to running: %s", err)
	}
	return nil
}
