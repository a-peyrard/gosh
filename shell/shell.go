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
	commandStore *command.Store
}

//goland:noinspection GoExportedFuncWithUnexportedType
func New() *gosh {
	return &gosh{
		commandStore: command.NewStore(),
	}
}

func (g *gosh) CommandStore() *command.Store {
	return g.commandStore
}

func (g *gosh) Run() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    g.commandStore.Completer(),
		InterruptPrompt: "^C",
		EOFPrompt:       "bye",

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
			continue
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		commandLine := command.ParseLine(line)

		exists, cmd := g.commandStore.Lookup(commandLine.Name)
		if !exists {
			log.Println("unknown command: ", commandLine.Name)
			continue
		}
		err = cmd.Executor()(commandLine, in, out)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("error executing command %s: %+v\n", commandLine.Name, err)
		}
	}
}

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}
