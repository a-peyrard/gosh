package main

import (
	"github.com/a-peyrard/gosh/builtins"
	"github.com/a-peyrard/gosh/shell"
	"sync"
)

func main() {
	gosh := shell.New()

	_ = gosh.AddCommand(builtins.PluginCommand())
	_ = gosh.AddCommand(builtins.EchoCommand())
	_ = gosh.AddCommand(builtins.PwdCommand())
	_ = gosh.AddCommand(builtins.ExitCommand())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		gosh.Run()
	}()

	wg.Wait()
}
