package main

import (
	"github.com/a-peyrard/gosh/builtins"
	"github.com/a-peyrard/gosh/shell"
	"sync"
)

func main() {
	gosh := shell.New()

	_ = gosh.CommandStore().AddCommand(builtins.PluginCommand(gosh.CommandStore()))
	_ = gosh.CommandStore().AddCommand(builtins.EchoCommand())
	_ = gosh.CommandStore().AddCommand(builtins.PwdCommand())
	_ = gosh.CommandStore().AddCommand(builtins.ExitCommand())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		gosh.Run()
	}()

	wg.Wait()
}
