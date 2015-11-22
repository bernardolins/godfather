package task

import (
	"fmt"
	"strings"
)

func Validate(task Task) {

}

func validateMetadata(m Metadata) bool {
	if len(m.Name) == 0 {
		fmt.Println("Field Name is required")
		return false
	}

	return true
}

func isValidName(name string) bool {
	invalidCharacters := []string{".", "/", "\\", "$", "%"}

	for _, invalidChar := range invalidCharacters {
		if strings.Contains(name, invalidChar) {
			fmt.Println("Name should not contain %q", invalidChar)
			return false
		}
	}

	return true
}
