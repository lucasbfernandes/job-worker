package storage

import (
	"log"
	"os"
)

const (
	logsDirPermission = 0700

	// This will create files inside pwd/logs
	defaultLogsDir = "logs"
)

func GetLogsDir() string {
	envLogsDIR, envExists := os.LookupEnv("LOGS_DIR")
	if envExists {
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
