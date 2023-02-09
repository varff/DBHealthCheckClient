package settings

import (
	"errors"
	"os"
)

func GetEnvDefault(key, defaultValue string) (string, error) {
	value := os.Getenv(key)
	if key == "" {
		if defaultValue == "" {
			return "", errors.New("environment variable isn't set")
		}
		return defaultValue, nil
	}

	return value, nil

}
