package storage

import (
	"fmt"
	"log"
	"os"
)

const (
	logsDirPermission = 0700
)

func CreateLogsDir() {
	logsDIR := os.Getenv("LOGS_DIR")

	if _, err := os.Stat(logsDIR); os.IsNotExist(err) {
		err = os.Mkdir(logsDIR, logsDirPermission)
		if err != nil {
			log.Fatalf(fmt.Sprintf("failed to create logs directory: %s\n", err))
		}
	}
}

func DeleteLogsDir() error {
	logsDIR := os.Getenv("LOGS_DIR")

	err := os.RemoveAll(logsDIR)
	if err != nil {
		log.Printf("failed to delete files inside logs dir: %s\n", err)
		return err
	}

	return nil
}