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

func (s *ServerInteractor) CreateJob(createJobRequest dto.CreateJobRequest) (*dto.CreateJobResponse, error) {
	job := createJobRequest.ToJob()

	process, err := s.createWorkerProcess(createJobRequest.Command)
	if err != nil {
		return nil, err
	}

	err = s.createWorkerProcessOutputFiles(process, job.ID)
	if err != nil {
		return nil, err
	}

	err = process.Start()
	if err != nil {
		return nil, err
	}

	savedJob, err := s.persistJob(job, process)
	if err != nil {
		return nil, err
	}

	go s.waitForExitReason(savedJob, process)

	return &dto.CreateJobResponse{ID: savedJob.ID}, nil
}

func (s *ServerInteractor) waitForExitReason(job *jobEntity.Job, process *worker.Process) {
	exitReason := <-process.ExitChannel
	var err error

	switch exitReason.ExitCode {
	case -1:
		err = s.finishJobWithStatusAndCode(*job, jobEntity.STOPPED, exitReason.ExitCode)
	case 0:
		err = s.finishJobWithStatusAndCode(*job, jobEntity.COMPLETED, exitReason.ExitCode)
	default:
		err = s.finishJobWithStatusAndCode(*job, jobEntity.FAILED, exitReason.ExitCode)
	}

	if err != nil {
		log.Printf("%s\n", err.Error())
	}
}

func (s *ServerInteractor) persistJob(job *jobEntity.Job, process *worker.Process) (*jobEntity.Job, error) {
	job.SetProcess(process)

	err := s.Database.UpsertJob(job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (s *ServerInteractor) createWorkerProcess(command []string) (*worker.Process, error) {
	process, err := worker.NewProcess(command)
	if err != nil {
		return nil, err
	}
	return process, nil
}

func (*ServerInteractor) createWorkerProcessOutputFiles(process *worker.Process, jobID string) error {
	outfile, err := repository.CreateLogFile(jobID)
	if err != nil {
		return err
	}
	process.SetStdoutWriter(outfile)
	process.SetStderrWriter(outfile)
	return nil
}

func (s *ServerInteractor) finishJobWithStatusAndCode(job jobEntity.Job, status string, exitCode int) error {
	finishedAt := time.Now()

	job.Status = status
	job.FinishedAt = &finishedAt
	job.ExitCode = exitCode

	err := s.Database.UpsertJob(&job)
	if err != nil {
		return fmt.Errorf("failed to update job with status %s: %s", status, err)
	}
	return nil
}
