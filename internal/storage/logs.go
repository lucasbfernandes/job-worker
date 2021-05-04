package storage

import (
	"log"
	"os"
)

const (
	logsDirPermission = 0700
)

func GetLogsDir() string {
	return os.Getenv("LOGS_DIR")
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
