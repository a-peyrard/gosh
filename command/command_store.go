package command

import (
	"github.com/chzyer/readline"
	"github.com/go-errors/errors"
)

type Store struct {
	commands map[string]Command
	version  int
}

func NewStore() *Store {
	return &Store{
		commands: map[string]Command{},
	}
}

func (s *Store) Lookup(commandAlias string) (found bool, cmd Command) {
	cmd, found = s.commands[commandAlias]
	return
}

func (s *Store) AddCommand(cmd Command) (err error) {
	return s.AddCommandWithAlias(cmd.Name(), cmd)
}

func (s *Store) AddCommandWithAlias(commandAlias string, cmd Command) (err error) {
	found, _ := s.Lookup(commandAlias)
	if found {
		err = errors.Errorf("unable to add command %s, alias is already in use", commandAlias)
	} else {
		s.commands[commandAlias] = cmd
		s.version++
	}

	return
}

func (s *Store) RemoveCommand(commandAlias string) {
	delete(s.commands, commandAlias)
	s.version++
}

func (s *Store) Completer() readline.PrefixCompleterInterface {
	// fixme: make it dynamic, we probably want to create a dynamic completer,
	// fixme: using a supplier with memoization, and the memo expire if the version of the store
	// fixme: has been changed
	return readline.NewPrefixCompleter(BuildCompleterForCommands(s.commands)...)
}
