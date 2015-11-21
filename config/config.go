package config

import (
	"fmt"
	"os"
)

func GetGodfatherPath() string {
	path := os.Getenv("GODFATHER_PATH")

	if len(path) == 0 {
		fmt.Println("$GODFATHER_PATH is not set")
		return ""
	}

	return path
}
