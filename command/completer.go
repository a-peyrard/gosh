package command

import (
	"github.com/chzyer/readline"
)

func BuildCompleterForCommands(commands map[string]Command) []readline.PrefixCompleterInterface {
	prefixCompleters := make([]readline.PrefixCompleterInterface, 0)
	for name, cmd := range commands {
		switch t := cmd.(type) {
		case WithCompleter:
			prefixCompleters = append(prefixCompleters, t.Completer(name))
		default:
			prefixCompleters = append(prefixCompleters, readline.PcItem(name))
		}
	}
	return prefixCompleters
}
