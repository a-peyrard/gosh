package builtins

import (
	"github.com/a-peyrard/gosh/command"
	"github.com/a-peyrard/gosh/shellio"
)

func PluginCommand() command.Command {
	return command.NewCommandWithSubsBuilder("plugin").
		Description("manage plugins").
		AddSubCommand(
			command.
				NewCommandBuilder("install").
				Description("install a plugin").
				Executor(func(cmd *command.Line, reader shellio.Reader, writer shellio.Writer) {
					writer.WriteLine("install plugin...")
				}).
				Build(),
		).
		AddSubCommand(
			command.
				NewCommandBuilder("uninstall").
				Description("uninstall a plugin").
				Executor(func(cmd *command.Line, reader shellio.Reader, writer shellio.Writer) {
					writer.WriteLine("uninstall plugin...")
				}).
				Build(),
		).
		AddSubCommand(
			command.
				NewCommandBuilder("list").
				Description("list installed plugins").
				Executor(func(cmd *command.Line, reader shellio.Reader, writer shellio.Writer) {
					writer.WriteLine("list plugin...")
				}).
				Build(),
		).
		Build()
}
