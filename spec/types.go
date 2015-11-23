package spec

import ()

type Task struct {
	Metadata Metadata          `yaml: "metadata"`
	Env      map[string]string `yaml: "env,omitempty"`
	Before   Command           `yaml: "before"`
	Command  Command           `yaml: "command"`
	After    Command           `yaml: "after"`
}

type Metadata struct {
	// Name is the name of the project
	Name string `yaml: "name"`
	// Selectors is aditional metadata to select groups of projects
	Selectors map[string]string `yaml: "selectors,omitempty"`
}

type Command struct {
	Exec      string   `yaml: "exec"`
	Arguments []string `yaml: "arguments"`
}
