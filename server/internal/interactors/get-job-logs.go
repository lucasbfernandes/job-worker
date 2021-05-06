package interactors

import (
	"io"
	"log"
	"os"
	"server/internal/repository"
)

func GetJobLogs(jobID string) (*string, error) {
	_, err := repository.GetJobOrFail(jobID)
	if err != nil {
		return nil, err
	}

	logFile, err := repository.GetLogFile(jobID)
	if err != nil {
		return nil, err
	}
	defer closeFile(logFile)

	logFileContent, err := getLogContent(logFile)
	if err != nil {
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
