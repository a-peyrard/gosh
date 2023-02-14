package shell

import (
	"github.com/a-peyrard/gosh/command"
	"github.com/a-peyrard/gosh/shellio"
	"github.com/chzyer/readline"
	"io"
	"log"
	"strings"
)

type gosh struct {
	commands map[string]*command.Command
}

//goland:noinspection GoExportedFuncWithUnexportedType
func New() *gosh {
	return &gosh{
		commands: make(map[string]*command.Command, 0),
	}
}

func (g *gosh) AddCommand(c *command.Command) error {
	// fixme validate the command correctness
	g.commands[c.Name] = c

	return nil
}

func (g *gosh) Run() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    buildCompleter(g.commands),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = l.Close()
	}()

	l.CaptureExitSignal()
	log.SetOutput(l.Stderr())

	in := shellio.NewDefaultReader()
	out := shellio.NewDefaultWriter()

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		commandLine := command.ParseLine(line)

		c, exists := g.commands[commandLine.Name]
		if !exists {
			log.Println("unknown command: ", commandLine.Name)
			continue
		}
		if c.ExecutorE != nil {
			err := c.ExecutorE(commandLine, in, out)
			if err != nil {
				log.Printf("error executing command %s: %+v\n", commandLine.Name, err)
			}
		} else {
			c.Executor(commandLine, in, out)
		}
	}
}

func buildCompleter(commands map[string]*command.Command) readline.AutoCompleter {
	prefixCompleters := make([]readline.PrefixCompleterInterface, 0)
	for name := range commands {
		prefixCompleters = append(prefixCompleters, readline.PcItem(name))
	}
	return readline.NewPrefixCompleter(prefixCompleters...)
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}
