package config

import (
	"fmt"
	"os"
)

func GetGodfatherPath() string {
	path := os.Getenv("GODFATHER_PATH")

	if len(path) == 0 {
		fmt.Println("$GODFATHER_PATH is not set")
		os.Exit(1)
	}

	return path
}
