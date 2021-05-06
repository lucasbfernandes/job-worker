package integration

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"server/internal/repository"
	"server/internal/storage"
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

	err = storage.DeleteLogsDir()
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
