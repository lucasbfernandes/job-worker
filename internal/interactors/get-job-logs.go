package interactors

import (
	"fmt"
	"io/ioutil"
	"job-worker/internal/repository"
	"log"
	"os"
)

// TODO Improve memory utilization. This is loading both files entirely in memory.
func GetJobLogs(jobID string) (string, error) {
	_, err := repository.GetJobOrFail(jobID)
	if err != nil {
		log.Printf("could not get job logs: %s\n", err)
		return "", err
	}

	stdoutFile, err := repository.GetStdoutLogFile(jobID)
	if err != nil {
		log.Printf("could not get stdout file: %s\n", err)
		return "", err
	}
	defer closeFile(stdoutFile)

	stderrFile, err := repository.GetStderrLogFile(jobID)
	if err != nil {
		log.Printf("could not get stderr file: %s\n", err)
		return "", err
	}
	defer closeFile(stderrFile)

	jobLogsResponse, err := getJobLogsResponse(stdoutFile, stderrFile)
	if err != nil {
		log.Printf("could not format get job logs response: %s\n", err)
		return "", err
	}

	return jobLogsResponse, nil
}

func getJobLogsResponse(stdoutFile *os.File, stderrFile *os.File) (string, error) {
	stdoutContent, err := ioutil.ReadAll(stdoutFile)
	if err != nil {
		log.Printf("could not get stdout content: %s\n", err)
		return "", err
	}
	stderrContent, err := ioutil.ReadAll(stderrFile)
	if err != nil {
		log.Printf("could not get stderr content: %s\n", err)
		return "", err
	}
	return fmt.Sprintf("stdout:\n%s\nstderr:\n%s", stdoutContent, stderrContent), nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Printf("failed to close file with error: %s\n", err)
	}
}
