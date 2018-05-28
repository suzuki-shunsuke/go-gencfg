package infra

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/mattn/go-shellwords"
)

type (
	// OSExecExecuter executes command.
	OSExecExecuter struct{}
)

// Exec executes a command.
func (executer OSExecExecuter) Exec(cmd, dest string) error {
	args, err := shellwords.Parse(cmd)
	if err != nil {
		return err
	}
	if len(args) == 0 {
		return nil
	}
	args = append(args, dest)
	fmt.Println(strings.Join(args, " "))
	return exec.Command(args[0], args[1:]...).Run()
}
