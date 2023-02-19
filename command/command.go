package command

import (
	"github.com/a-peyrard/gosh/shellio"
)

type Command interface {
	Name() string
	Description() string
	Executor() func(cmd *Line, reader shellio.Reader, writer shellio.Writer)
	UnsafeExecutor() func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error
}

type baseCommand struct {
	name           string
	description    string
	executor       func(cmd *Line, reader shellio.Reader, writer shellio.Writer)
	unsafeExecutor func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error
}

func (c *baseCommand) Name() string {
	return c.name
}

func (c *baseCommand) Description() string {
	return c.description
}

func (c *baseCommand) Executor() func(cmd *Line, reader shellio.Reader, writer shellio.Writer) {
	return c.executor
}

func (c *baseCommand) UnsafeExecutor() func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error {
	return c.unsafeExecutor
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
	b.executor = executor

	return b
}

func (b *baseCommandBuilder) UnsafeExecutor(executor func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error) *baseCommandBuilder {
	b.unsafeExecutor = executor

	return b
}

func (b *baseCommandBuilder) Build() Command {
	if b.executor == nil && b.unsafeExecutor == nil {
		panic("at least an executor or an unsafeExecutor must be defined")
	}
	return &b.baseCommand
}
