package base

import (
	"strings"
)

type StringList []string

func (l *StringList) String() string {
	return strings.Join(*l, ", ")
}

func (l *StringList) Set(value string) error {
	*l = append(*l, value)
	return nil
}

func (l *StringList) Type() string {
	return "StringList"
}
