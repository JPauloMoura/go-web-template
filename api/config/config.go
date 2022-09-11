package config

import "os"

const _DEFAULT_PORT = "3002"

// GetPort returns the port where the server should run
func GetPort() string {
	switch port, exist := os.LookupEnv("PORT"); {
	case exist:
		return port
	default:
		return _DEFAULT_PORT
	}
}
