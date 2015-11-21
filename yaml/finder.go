package yaml

import (
	"fmt"
	"github.com/bernardolins/godfather/config"
	"os"
	"path/filepath"
	"strings"
)

//Finds YAML file on given directory
func Find() {
	err := filepath.Walk(config.GetGodfatherPath(), visit)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && isGodfatherFile(f.Name()) {
		fmt.Println(path)
	}
	return nil
}

func isGodfatherFile(filename string) bool {
	if strings.Contains(filename, ".godfather") && isYaml(filename) {
		return true
	}

	return false
}

func isYaml(filename string) bool {
	if strings.Contains(filename, ".yml") || strings.Contains(filename, ".yaml") {
		return true
	}

	return false
}
