package interactors

import (
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

	logFile, err := repository.GetLogFile(jobID)
	if err != nil {
		log.Printf("could not get stdout file: %s\n", err)
		return "", err
	}
	defer closeFile(logFile)

	logFileContent, err := ioutil.ReadAll(logFile)
	if err != nil {
		log.Printf("could not get log file content: %s\n", err)
		return "", err
	}

	return string(logFileContent), nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Printf("failed to close file with error: %s\n", err)
	}
}
