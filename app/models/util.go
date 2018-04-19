package models

import (
	"os"
)

// getEnvWithDefault ...
func getEnvWithDefault(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
