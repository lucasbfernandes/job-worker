package interactors

import (
	"job-worker/internal/dto"
	"job-worker/pkg/worker"
	"log"
	"os"
	"time"
)

func CreateJob(createJobRequest dto.CreateJobRequest) (string, error) {
	stdout, err := os.Create("./out/stdout.txt")
	if err != nil {
		panic(err)
	}
	stderr, err := os.Create("./out/stderr.txt")
	if err != nil {
		panic(err)
	}

	// create process with lib
	process, err := worker.NewProcess(createJobRequest.Command, time.Duration(createJobRequest.TimeoutInSeconds))
	if err != nil {
		log.Printf("could not create process %s\n", err)
		return "", err
	}
	process.SetStdoutWriter(stdout)
	process.SetStderrWriter(stderr)

	err = process.Start()
	if err != nil {
		log.Printf("could not start process %s\n", err)
		return "", err
	}
	return "", nil
}
