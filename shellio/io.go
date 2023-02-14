package shellio

import "fmt"

type Reader interface {
	Ask(question string) string
}

type defaultReader struct{}

func NewDefaultReader() Reader {
	return &defaultReader{}
}

func (d *defaultReader) Ask(question string) string {
	return "something" // fixme
}

type Writer interface {
	WriteLine(line string)
}

type defaultWriter struct{}

func NewDefaultWriter() Writer {
	return &defaultWriter{}
}

func (d *defaultWriter) WriteLine(line string) {
	fmt.Println(line)
}
