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

func SetTestJWTCertPath() error {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	err := os.Setenv("JWT_CERT_FILE_PATH", path.Join(basepath, "fixtures/cert/mock-jwt.crt"))
	if err != nil {
		return err
	}
	return nil
}

func GetAdminUserMockAPIToken() string {
	return "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcGlUb2tlbiI6InFUTWFZSWZ3OHEzZXNaNkR2MnJRIn0.wUzwsMxeKkZrt8uct1eeFlVPh8JejA6TlOQkALZ9I7Z0hWJkAUlJgQJEaM-p2NMx2iGVJI81ueXciN6BQNI8TCgdKXgfJqFbIeE-pbPyIpHBXp_7BVAUIHsLvoLsQbA0o_Lxw0uO21mbqsG-odd3zOMFmGb7u1lWN2lh8UqnOOUbPmSL4N4FSu5A6LG5eN_0OAIQySxKKrcGafPiikDfKLzSZQHcl_gkOJZrKkAIoUgAEluL4djFtulIxtfzWU8N_nWc5zcUCOnT6VEPXItxbpiRXRromWPfu3oIq6mHm-6qW_5RBg9UXU0m_FwpyCHCovvVPCCLaW7vX00AKS26Ew"
}

func GetUserMockAPIToken() string {
	return "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcGlUb2tlbiI6IjlFekdKT1RjTUhGTVhwaGZ2QXVNIn0.fMOIjg6XVd3hg_WumX9yjbRIgRUvwfiWIkqP5b6w83OgglIRxOIesey-ZxTg8o_tu1ByQPDnGOFtK577gc4fx2wulqYx9_wDgrLFgtAz02J1YZnhCC9jbhbX1dOc8lhpPba3Yd0GWYa-ZT7QK24mZvKyccfUSRYj819WMyW6beIbHLjciRkN1MOBjemtDWFNix4AdnaJf1iIj6ffXk-2PvUEWFWAhpPUgD-OU52Vdz2TOWF_xikN86F6sawXoLhrfLq3LNoUuIEjkpaEInq0loHxyN4yT72pncg-JUt_3i7p_6ShIrIjdAYGJLN58BEDcJoV-p7b2xnNZr0ixVdZLg"
}

func GetNonExistentUserAPIToken() string {
	return "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcGlUb2tlbiI6Indyb25nLXRva2VuIn0.FAsTw6XfK7eN-c0EBC11GjQ5YyGkD0uThFL1ERPm8WmJiNUjmo71Aq9lQlWIdJooRLr80ZfQPoSLA7jQbiOFUM5Nw2gTcQ9zBqZALkcE89TnRZP6hp7-lwoPFZI9-gXafaUSj8MmfPoPnMP4RPExvrRTzyY2y4f94ZO9jLEbi-H5Ev5tWLOIl8PY6dMvS9zgMSDva0ZweBdDK8ksJLykegfucLShOcqdejQlrcYIxBAq7hpr8KQhCobF3SFA8krJLt_9BH4hffdLJkyhfI-yY9Q7kyeHUp4mvuvUOyHAIIDDxuQg2ZcyzK8ByIM_4xVFKdRbTQwSij-jDJZ5EGXq-Q"
}

func GetInvalidBearerToken() string {
	return "..FAsTwEv5tWLOIl8Xq-Q"
}
