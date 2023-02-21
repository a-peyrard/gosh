package command

import (
	"fmt"
	"github.com/a-peyrard/gosh/shellio"
	"github.com/chzyer/readline"
)

type baseCommandWithSubs struct {
	subCommands map[string]Command
	baseCommand
}

func (b *baseCommandWithSubs) Completer(commandAlias string) readline.PrefixCompleterInterface {
	return readline.PcItem(
		commandAlias,
		BuildCompleterForCommands(b.subCommands)...,
	)
}

func (b *baseCommandWithSubs) Executor() Executor {
	return func(cmd *Line, reader shellio.Reader, writer shellio.Writer) (err error) {
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
		err = subCommand.Executor()(cmd, reader, writer)

		return
	}
}

type baseCommandWithSubsBuilder struct {
	baseCommandWithSubs
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewCommandWithSubsBuilder(name string) *baseCommandWithSubsBuilder {
	return &baseCommandWithSubsBuilder{
		baseCommandWithSubs: baseCommandWithSubs{
			subCommands: map[string]Command{},
			baseCommand: baseCommand{name: name},
		},
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
	return &b.baseCommandWithSubs
}
