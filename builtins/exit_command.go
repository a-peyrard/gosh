package builtins

import (
	"github.com/a-peyrard/gosh/command"
	"github.com/a-peyrard/gosh/shellio"
	"io"
)

func ExitCommand() command.Command {
	return command.
		NewCommandBuilder("exit").
		Description("exit the current shell").
		UnsafeExecutor(func(cmd *command.Line, reader shellio.Reader, writer shellio.Writer) error {
			return io.EOF
		}).
		Build()
}
