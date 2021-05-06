package interactors

const (
	jobsPath         = "/jobs"
	stopJobsPath     = "/stop"
	getJobStatusPath = "/status"
	getJobLogsPath   = "/logs"

	dateLayout = "2006-01-02 15:04:05.000"
)

type WorkerCLIInteractor struct{}

func NewWorkerCLIInteractor() *WorkerCLIInteractor {
	return &WorkerCLIInteractor{}
}
