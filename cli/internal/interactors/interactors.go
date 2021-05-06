package interactors

const (
	jobsPath         = "/jobs"
	stopJobsPath     = "/stop"
	getJobStatusPath = "/status"
	getJobLogsPath   = "/logs"
)

type WorkerCLIInteractor struct{}

func NewWorkerCLIInteractor() *WorkerCLIInteractor {
	return &WorkerCLIInteractor{}
}
