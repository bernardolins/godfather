package config

import (
	"fmt"
	"os"
)

func GetGodfatherPath() string {
	path := os.Getenv("GODFATHER_PATH")

	if len(path) == 0 {
		fmt.Println("Warning: $GODFATHER_PATH is not set, assuming path as $PWD instead")
		return os.Getenv("PWD")
	}

	return path
}
