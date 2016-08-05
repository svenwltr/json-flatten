package main

import (
	"fmt"
	"io"
	"strings"
)

type Assignments []Assignment

func (a Assignments) PrintTo(w io.Writer) {
	for _, assignment := range a {
		assignment.PrintTo(w)
		fmt.Fprint(w, "\n")
	}
}

type Assignment struct {
	Path  []Key
	Value interface{}
}

func (a Assignment) PrintTo(w io.Writer) {
	keys := make([]string, len(a.Path))
	for i, key := range a.Path {
		keys[i] = key.String()
	}
	fmt.Fprintf(w, "%s: %v",
		strings.Join(keys, "."),
		a.Value)
}

type Key interface {
	String() string
}

type ObjectName string

func (o ObjectName) String() string {
	return string(o)
}

type ArrayIndex int

func (i ArrayIndex) String() string {
	return fmt.Sprintf("%d", i)
}

func Flatten(data interface{}) Assignments {
	return flatten([]Key{}, data)
}

func flatten(path []Key, value interface{}) Assignments {
	switch value := value.(type) {
	case map[string]interface{}:
		return flattenMap(path, value)
	case []interface{}:
		return flattenArray(path, value)
	default:
		return Assignments{Assignment{path, value}}
		//panic(fmt.Sprintf("Unknown type: %#v", value))
	}
}

func flattenMap(path []Key, value map[string]interface{}) Assignments {
	out := make(Assignments, 0)
	for k, v := range value {
		p := append(path, ObjectName(k))
		out = append(out, flatten(p, v)...)
	}

	return out
}

func flattenArray(path []Key, value []interface{}) Assignments {
	out := make(Assignments, 0)
	for i, v := range value {
		p := append(path, ArrayIndex(i))
		out = append(out, flatten(p, v)...)
	}

	return out

}
