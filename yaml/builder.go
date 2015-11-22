package yaml

import (
	"fmt"
	"github.com/bernardolins/godfather/task"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func BuildTask(path string) task.Task {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var t task.Task
	err = yaml.Unmarshal(f, &t)

	if !task.Validate(t) {
		fmt.Println("An error occurred when creating task at", path)
	}

	if err != nil {
		panic(err)
	}

	return t
}
