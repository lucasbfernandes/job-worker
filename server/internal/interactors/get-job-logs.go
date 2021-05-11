package interactors

import (
	"io"
	"log"
	"os"
	"server/internal/repository"
	"strings"
)

func (s *ServerInteractor) GetJobLogs(jobID string) (*string, error) {
	_, err := s.Database.GetJobOrFail(jobID)
	if err != nil {
		return nil, err
	}

	logFile, err := repository.GetLogFile(jobID)
	if err != nil {
		return nil, err
	}
	defer closeFile(logFile)

	logFileContent, err := s.getLogContent(logFile)
	if err != nil {
		return nil, err
	}

	return logFileContent, nil
}

func (s *ServerInteractor) getLogContent(logFile *os.File) (*string, error) {
	buf := new(strings.Builder)

	_, err := io.Copy(buf, logFile)
	if err != nil {
		return nil, err
	}

	logContent := buf.String()
	return &logContent, nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Printf("failed to close file with error: %s\n", err)
	}
}
