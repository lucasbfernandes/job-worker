package repository

import (
	"fmt"
	"os"
	"path"
	"server/internal/storage"
)

func CreateLogFile(jobID string) (*os.File, error) {
	if jobID == "" {
		return nil, fmt.Errorf("empty jobID")
	}
	logsDIR := storage.GetLogsDir()
	logFile, err := os.Create(path.Join(logsDIR, jobID))
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func GetLogFile(jobID string) (*os.File, error) {
	if jobID == "" {
		return nil, fmt.Errorf("empty jobID")
	}
	logsDIR := storage.GetLogsDir()
	logFile, err := os.Open(path.Join(logsDIR, jobID))
	if err != nil {
		return nil, err
	}
	return logFile, nil
}