package exec

import (
	"github.com/bernardolins/godfather/command"
	"github.com/bernardolins/godfather/spec"
)

func Exec(task spec.Task) {
	createEnvironment(task)
	execCommand(task.Command)
}

func createEnvironment(task spec.Task) {
	if len(task.Env) > 0 {
		for variable, value := range task.Env {
			command.SetupEnvironmentVariable(variable, value)
		}
	}
}

func execBeforeHook(cmd spec.Command) {

}

func execAfterHook(cmd spec.Command) {

}

func execCommand(cmd spec.Command) {
	command.Exec(cmd)
}
