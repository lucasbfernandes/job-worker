package repository

import (
	"job-worker/internal/storage"
	"os"
	"path"
)

func CreateStdoutLogFile(jobID string) (*os.File, error) {
	logsDIR := storage.GetLogsDir()
	stdoutFile, err := os.Create(path.Join(logsDIR, jobID+"-stdout"))
	if err != nil {
		return nil, err
	}
	return stdoutFile, nil
}

func CreateStderrLogFile(jobID string) (*os.File, error) {
	logsDIR := storage.GetLogsDir()
	stderrFile, err := os.Create(path.Join(logsDIR, jobID+"-stderr"))
	if err != nil {
		return nil, err
	}
	return stderrFile, nil
}

func GetStdoutLogFile(jobID string) (*os.File, error) {
	logsDIR := storage.GetLogsDir()
	stdoutFile, err := os.Open(path.Join(logsDIR, jobID+"-stdout"))
	if err != nil {
		return nil, err
	}
	return stdoutFile, nil
}

func GetStderrLogFile(jobID string) (*os.File, error) {
	logsDIR := storage.GetLogsDir()
	stderrFile, err := os.Open(path.Join(logsDIR, jobID+"-stderr"))
	if err != nil {
		return nil, err
	}
	return stderrFile, nil
}
