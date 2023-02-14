package command

import (
	"github.com/a-peyrard/gosh/shellio"
)

//type Command interface {
//	Name() string
//	Description() string
//	Executor() func(cmd *Line, reader shell.Reader, writer shell.Writer)
//}

type Command struct {
	Name        string
	Description string
	Executor    func(cmd *Line, reader shellio.Reader, writer shellio.Writer)
	ExecutorE   func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error
}
