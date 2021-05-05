package config

import "os"

const (
	defaultServerURL = "http://localhost:8883"
)

func GetDefaultServerURL() string {
	envDefaultServer, envExists := os.LookupEnv("DEFAULT_SERVER")
	if envExists && envDefaultServer != "" {
		return envDefaultServer
	}
	return defaultServerURL
}
