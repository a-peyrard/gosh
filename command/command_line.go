package command

import "strings"

type Line struct {
	Name      string
	Arguments []string
}

func ParseLine(line string) *Line {
	tokens := strings.Split(line, " ")
	return &Line{
		Name:      tokens[0],
		Arguments: tokens[1:],
	}
}
