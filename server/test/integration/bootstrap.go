package integration

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"server/internal/repository"
)

func BootstrapTestEnvironment() error {
	err := setTestLogsDirEnv()
	if err != nil {
		return err
	}
	return nil
}

func RollbackState(db *repository.InMemoryDatabase) error {
	err := db.DeleteAllJobs()
	if err != nil {
		return err
	}

	err = repository.DeleteLogsDir()
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
	files, err := ioutil.ReadDir(repository.GetLogsDir())
	if err != nil {
		return nil, err
	}
	filesNumber := len(files)
	return &filesNumber, nil
}
