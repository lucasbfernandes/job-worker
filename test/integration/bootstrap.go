package integration

import (
	"job-worker/internal/repository"
	"job-worker/internal/storage"
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
	storage.CreateDB()
	return nil
}

func RollbackState() error {
	err := repository.DeleteAllJobs()
	if err != nil {
		return err
	}

	err = storage.DeleteLogsDir()
	if err != nil {
		return err
	}

	return nil
}

func setTestLogsDirEnv() error {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)

	err := os.Setenv("LOGS_DIR", path.Join(basepath, "logs-test"))
	if err != nil {
		return err
	}
	return nil
}
