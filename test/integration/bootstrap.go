package integration

import (
	"io/ioutil"
	"job-worker/internal/repository"
	"job-worker/internal/storage"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func BootstrapTestEnvironment() error {
	err := setTestLogsDirEnv()
	if err != nil {
		return err
	}
	err = storage.CreateDB()
	if err != nil {
		return err
	}
	return nil
}

func RollbackState() error {
	err := repository.DeleteAllJobs()
	if err != nil {
		return err
	}

	err = deleteLogsDir()
	if err != nil {
		return err
	}

	return nil
}

func setTestLogsDirEnv() error {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := os.Setenv("LOGS_DIR", path.Join(basepath, "logs-test"))
	if err != nil {
		return err
	}
	return nil
}

func GetNumberOfLogFiles() (*int, error) {
	files, err := ioutil.ReadDir(storage.GetLogsDir())
	if err != nil {
		return nil, err
	}
	filesNumber := len(files)
	return &filesNumber, nil
}

func deleteLogsDir() error {
	logsDIR := storage.GetLogsDir()

	err := os.RemoveAll(logsDIR)
	if err != nil {
		log.Printf("failed to delete files inside logs dir: %s\n", err)
		return err
	}

	return nil
}
