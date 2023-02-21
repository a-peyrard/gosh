package command

import "github.com/a-peyrard/gosh/shellio"

type baseCommand struct {
	name        string
	description string
	executor    func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error
}

func (c *baseCommand) Name() string {
	return c.name
}

func (c *baseCommand) Description() string {
	return c.description
}

func (c *baseCommand) Executor() func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error {
	return c.executor
}

type baseCommandBuilder struct {
	baseCommand
}

//goland:noinspection GoExportedFuncWithUnexportedType
func NewCommandBuilder(name string) *baseCommandBuilder {
	return &baseCommandBuilder{baseCommand{name: name}}
}

func (b *baseCommandBuilder) Description(description string) *baseCommandBuilder {
	b.description = description

	return b
}

func (b *baseCommandBuilder) Executor(executor func(cmd *Line, reader shellio.Reader, writer shellio.Writer)) *baseCommandBuilder {
	b.executor = func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error {
		executor(cmd, reader, writer)
		return nil
	}

	return b
}

func (b *baseCommandBuilder) UnsafeExecutor(executor func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error) *baseCommandBuilder {
	b.executor = executor

	return b
}

func (b *baseCommandBuilder) Build() Command {
	if b.executor == nil {
		panic("an executor must be defined!")
	}
	return &b.baseCommand
}
