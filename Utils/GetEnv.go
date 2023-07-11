package util

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		fmt.Printf("Could get Env %s\n", key)
		return ""
	}
	return env
}
