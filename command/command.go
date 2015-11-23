package command

import (
	"bufio"
	"github.com/bernardolins/godfather/spec"
	"log"
	"os"
	"os/exec"
	"strings"
)

func SetupEnvironmentVariable(name string, value string) {
	os.Setenv(name, value)
}

func Exec(cmd spec.Command) {
	args := strings.Join(cmd.Arguments, " ")
	command := exec.Command(cmd.Exec, args)

	stdout, err := command.StdoutPipe()

	if err != nil {
		panic(err)
	}

	if err := command.Start(); err != nil {
		panic(err)
	}

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Printf(in.Text())
	}

	if err := in.Err(); err != nil {
		panic(err)
	}
}
