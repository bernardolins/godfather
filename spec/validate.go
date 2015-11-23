package spec

import (
	"fmt"
	"strings"
)

func Validate(task Task) bool {
	return validateMetadata(task.Metadata)
}

func validateMetadata(m Metadata) bool {
	if len(m.Name) == 0 {
		fmt.Println("Error: Field Name is required")
		return false
	}

	if !isValidName(m.Name) {
		return false
	}

	return true
}

func isValidName(name string) bool {
	invalidCharacters := []string{".", "/", "\\", "$", "%"}

	for _, invalidChar := range invalidCharacters {
		if strings.Contains(name, invalidChar) {
			fmt.Println("Error: Name should not contain", invalidChar)
			return false
		}
	}

	return true
}
