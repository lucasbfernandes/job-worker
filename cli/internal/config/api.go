package config

import "os"

const (
	DefaultServerURL = "http://localhost:8883"
)

func GetDefaultServerURL() string {
	envDefaultServer, envExists := os.LookupEnv("DEFAULT_SERVER_URL")
	if envExists && envDefaultServer != "" {
		return envDefaultServer
	}
	return DefaultServerURL
}
