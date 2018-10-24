package util

import (
	"os"
	"strconv"
)

// Env returns a string from the ENV, or fallback variable
func Env(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return fallback
}

// EnvInt returns an int from the ENV, or fallback variable
func EnvInt(key string, fallback int) int {
	if i, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return i
	}

	return fallback
}
