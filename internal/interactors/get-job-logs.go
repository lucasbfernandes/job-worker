package interactors

import (
	"io"
	"job-worker/internal/repository"
	"log"
	"os"
)

func GetJobLogs(jobID string) (*string, error) {
	_, err := repository.GetJobOrFail(jobID)
	if err != nil {
		log.Printf("could not get job logs: %s\n", err)
		return nil, err
	}

	logFile, err := repository.GetLogFile(jobID)
	if err != nil {
		log.Printf("could not get stdout file: %s\n", err)
		return nil, err
	}
	defer closeFile(logFile)

	logFileContent, err := getLogContent(logFile)
	if err != nil {
		log.Printf("could not get log file content: %s\n", err)
		return nil, err
	}

	return logFileContent, nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Printf("failed to close file with error: %s\n", err)
	}
}

func getLogContent(logFile *os.File) (*string, error) {
	bufferSize := 100
	contentBytes := make([]byte, 0)
	buffer := make([]byte, bufferSize)

	for {
		numberOfBytes, err := logFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		contentBytes = append(contentBytes, buffer[:numberOfBytes]...)
	}

	logContent := string(contentBytes)
	return &logContent, nil
}
