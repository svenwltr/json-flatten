package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Assignments []*Assignment

func (a Assignments) String() string {
	parts := make([]string, len(a))
	for i, assignment := range a {
		parts[i] = assignment.String()
	}
	return strings.Join(parts, "\n")
}

type Path []Name

func (p Path) String() string {
	out := "."
	if len(p) > 0 {
		buf := new(bytes.Buffer)
		for i, name := range p {
			_, isArray := name.(ArrayIndex)
			if i != 0 && !isArray {
				fmt.Fprint(buf, ".")
			}
			escaped := fmt.Sprintf("%v", name)
			escaped = strings.Replace(escaped, `\`, `\\`, -1)
			escaped = strings.Replace(escaped, `.`, `\.`, -1)
			escaped = strings.Replace(escaped, `=`, `\=`, -1)
			fmt.Fprintf(buf, "%s", escaped)
		}
		out = buf.String()
	}
	return out
}

type Name interface {
	String() string
}

type ObjectName string

func (n ObjectName) String() string {
	return string(n)
}

type ArrayIndex int

func (i ArrayIndex) String() string {
	return fmt.Sprintf("[%d]", i)
}

type Assignment struct {
	Path  Path
	Value Value
}

func (a Assignment) String() string {
	return fmt.Sprintf("%v=%v", a.Path, a.Value)
}

type Value interface {
}

type Raw string

func (r Raw) MarshalJSON() ([]byte, error) {
	return []byte(r), nil
}
