package command

import (
	"github.com/a-peyrard/gosh/shellio"
)

type Command interface {
	Name() string
	Description() string
	Executor() func(cmd *Line, reader shellio.Reader, writer shellio.Writer) error
}
