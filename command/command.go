package command

import (
	"github.com/a-peyrard/gosh/shellio"
	"github.com/chzyer/readline"
)

type (
	Executor = func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error

	Command interface {
		Name() string
		Description() string
		Executor() Executor
	}

	WithCompleter interface {
		Completer(commandAlias string) readline.PrefixCompleterInterface
	}
)
