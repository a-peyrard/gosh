package builtins

import (
	"github.com/a-peyrard/gosh/command"
	"github.com/a-peyrard/gosh/shellio"
	"strings"
)

func EchoCommand() *command.Command {
	return &command.Command{
		Name:        "echo",
		Description: "just echo something",
		Executor: func(cmd *command.Line, reader shellio.Reader, writer shellio.Writer) {
			writer.WriteLine(strings.Join(cmd.Arguments, " "))
		},
	}
}
