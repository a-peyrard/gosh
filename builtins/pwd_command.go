package builtins

import (
	"github.com/a-peyrard/gosh/command"
	"github.com/a-peyrard/gosh/shellio"
	"os"
)

func PwdCommand() command.Command {
	return command.NewCommandBuilder("pwd").
		Description("Print working directory").
		ExecutorE(func(cmd *command.Line, reader shellio.Reader, writer shellio.Writer) error {
			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			writer.WriteLine(dir)

			return nil
		}).
		Build()
}
