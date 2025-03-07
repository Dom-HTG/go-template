package env

import (
	"os"
	"strconv"
)

// { returns string }
func GetStringEnv(envName, fallback string) string {
	variable, ok := os.LookupEnv(envName)
	if !ok {
		return fallback
	}
	return variable
}

// { returns Int }
func GetIntEnv(envName string, fallback int) int {
	variable, ok := os.LookupEnv(envName)
	if !ok {
		return fallback
	}

	intEnv, err := strconv.Atoi(variable)
	if err != nil {
		return fallback
	}

	return intEnv
}
