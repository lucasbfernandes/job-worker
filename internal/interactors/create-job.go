package interactors

import (
	"job-worker/internal/dto"
	jobEntity "job-worker/internal/models/job"
	"job-worker/internal/repository"
	"job-worker/pkg/worker"
	"log"
	"os"
	"time"
)

func CreateJob(createJobRequest dto.CreateJobRequest) (dto.CreateJobResponse, error) {
	job, err := persistJob(createJobRequest)
	if err != nil {
		log.Printf("could not persist job %s\n", err)
		return dto.CreateJobResponse{}, err
	}

	process, err := createWorkerProcess(createJobRequest.Command, time.Duration(createJobRequest.TimeoutInSeconds))
	if err != nil {
		log.Printf("could not create process %s\n", err)
		setJobStatus(job, jobEntity.FAILED)
		return dto.CreateJobResponse{}, err
	}

	err = process.Start()
	if err != nil {
		log.Printf("could not start process %s\n", err)
		setJobStatus(job, jobEntity.FAILED)
		return dto.CreateJobResponse{}, err
	}

	// TODO trigger goroutine that will watch process exit reason channel
	setJobStatus(job, jobEntity.RUNNING)
	return dto.CreateJobResponse{ID: job.ID}, nil
}

func persistJob(request dto.CreateJobRequest) (jobEntity.Job, error) {
	job := request.ToJob()
	err := repository.CreateJob(job)
	if err != nil {
		return jobEntity.Job{}, err
	}
	return job, nil
}

func createWorkerProcess(command []string, timeoutInSeconds time.Duration) (worker.Process, error) {
	process, err := worker.NewProcess(command, timeoutInSeconds)
	if err != nil {
		return worker.Process{}, err
	}
	err = setWorkerProcessOutputFiles(process)
	if err != nil {
		return worker.Process{}, err
	}
	return process, nil
}

func setWorkerProcessOutputFiles(process worker.Process) error {
	// TODO create logs folder if it does not exist
	stdout, err := os.Create("./logs/stdout.txt")
	if err != nil {
		return err
	}
	stderr, err := os.Create("./logs/stderr.txt")
	if err != nil {
		return err
	}
	process.SetStdoutWriter(stdout)
	process.SetStderrWriter(stderr)
	return nil
}

func setJobStatus(job jobEntity.Job, status string) {
}
