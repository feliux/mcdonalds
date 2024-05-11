package commons

import "syscall"

// EnvString load an environment variable named 'key'.
// If the environment variable does not exist then return a 'fallback' as default value.
func EnvString(key, fallback string) string {
	if val, ok := syscall.Getenv(key); ok {
		return val
	}
	return fallback
}
