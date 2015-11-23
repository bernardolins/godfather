package yaml

import (
	"fmt"
	"github.com/bernardolins/godfather/spec"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func BuildTask(path string) spec.Task {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var t spec.Task
	err = yaml.Unmarshal(f, &t)

	if !spec.Validate(t) {
		fmt.Println("An error occurred when creating task at", path)
	}

	if err != nil {
		panic(err)
	}

	return t
}
