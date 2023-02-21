package command

import (
	"fmt"
	"github.com/a-peyrard/gosh/shellio"
)

type baseCommandWithSubsBuilder struct {
	subCommands map[string]Command
	baseCommand
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewCommandWithSubsBuilder(name string) *baseCommandWithSubsBuilder {
	return &baseCommandWithSubsBuilder{
		subCommands: map[string]Command{},
		baseCommand: baseCommand{name: name},
	}
}

func (b *baseCommandWithSubsBuilder) Description(description string) *baseCommandWithSubsBuilder {
	b.description = description

	return b
}

func (b *baseCommandWithSubsBuilder) AddSubCommand(cmd Command) *baseCommandWithSubsBuilder {
	_, exists := b.subCommands[cmd.Name()]
	if exists {
		panic(fmt.Sprintf(
			"command %s is already having a sub command named %s",
			b.baseCommand.name, cmd.Name(),
		))
	}
	b.subCommands[cmd.Name()] = cmd

	return b
}

func (b *baseCommandWithSubsBuilder) Build() Command {
	b.unsafeExecutor = func(cmd *Line, reader shellio.Reader, writer shellio.Writer) (err error) {
		if len(cmd.Arguments) == 0 {
			writer.WriteLine("display usage") // fixme
			return
		}
		subCommandName := cmd.Arguments[0]
		subCommand, exists := b.subCommands[subCommandName]
		if !exists {
			// fixme write as error! Should we return an error?
			writer.WriteLine(fmt.Sprintf("unknown sub command %s", subCommandName))
			return
		}

		cmd.Arguments = cmd.Arguments[1:]
		// fixme: shall we have only one executor and the executor not throwing errors is not in the
		// fixme: interface, but just a simplification of the builder?
		if subCommand.UnsafeExecutor() != nil {
			err = subCommand.UnsafeExecutor()(cmd, reader, writer)
		} else {
			subCommand.Executor()(cmd, reader, writer)
		}

		return
	}
	return &b.baseCommand
}
