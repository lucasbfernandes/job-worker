package integration

import (
	"cli/internal/config"
	"os"
)

func GetDefaultTestsServerURL() string {
	envDefaultServer, envExists := os.LookupEnv("DEFAULT_SERVER_MOCK_URL")
	if envExists && envDefaultServer != "" {
		return envDefaultServer
	}
	return config.DefaultServerURL
}
