package yaml

import (
	"github.com/bernardolins/godfather/task"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func BuildTask(path string) task.Task {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var task task.Task
	err = yaml.Unmarshal(f, &task)

	if err != nil {
		panic(err)
	}

	return task
}
