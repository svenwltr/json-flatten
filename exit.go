package main

import "fmt"

type ExitHandler func(code int)

const (
	ExitOk    = iota
	ExitError = iota
	ExitUsage = iota
)

func (app *Application) AssertNoError(err error) {
	if err != nil {
		fmt.Fprintln(app.Stderr, err)
		app.Exit(ExitError)
	}
}
