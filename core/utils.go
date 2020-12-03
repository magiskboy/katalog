package core

import "os"

// Getenv get env by key, return fallback if it doesn't exist
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
