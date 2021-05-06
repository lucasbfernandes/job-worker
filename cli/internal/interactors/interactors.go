package interactors

const (
	jobsPath         = "/jobs"
	stopJobsPath     = "/stop"
	getJobStatusPath = "/status"
)

type WorkerCLIInteractor struct{}

func NewWorkerCLIInteractor() *WorkerCLIInteractor {
	return &WorkerCLIInteractor{}
}
