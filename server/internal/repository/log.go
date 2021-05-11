package repository

import (
	"fmt"
	"log"
	"os"
	"path"
)

const (
	logsDirPermission = 0700

	// This will create files inside pwd/logs
	defaultLogsDir = "logs"
)

func GetLogsDir() string {
	envLogsDIR, envExists := os.LookupEnv("LOGS_DIR")
	if envExists && envLogsDIR != "" {
		return envLogsDIR
	}
	return defaultLogsDir
}

func CreateLogsDir() error {
	logsDIR := GetLogsDir()
	err := os.MkdirAll(logsDIR, logsDirPermission)
	if err != nil {
		log.Printf("failed to create logs directory: %s\n", err)
		return err
	}
	return nil
}

func DeleteLogsDir() error {
	logsDIR := GetLogsDir()

	err := os.RemoveAll(logsDIR)
	if err != nil {
		log.Printf("failed to delete logs dir: %s\n", err)
		return err
	}

	return nil
}

func CreateLogFile(jobID string) (*os.File, error) {
	if jobID == "" {
		return nil, fmt.Errorf("empty jobID")
	}
	logsDIR := GetLogsDir()
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
	logsDIR := GetLogsDir()
	logFile, err := os.Open(path.Join(logsDIR, jobID))
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
