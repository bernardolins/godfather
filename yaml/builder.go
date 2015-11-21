package yaml

import (
	"github.com/bernardolins/godfather/task"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func BuildTask(path string) task.T {
	f, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var task task.T
	err = yaml.Unmarshal(f, &task)

	if err != nil {
		log.Fatal(err)
	}

	return task
}
