package main

import (
	"fmt"
	"io"
	"os"
)

var (
	version = "unknown" // set by makefile
)

type Application struct {
	Args []string
	Exit ExitHandler

	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
}

func (app *Application) Run() {
	config := app.GetConfig()

	if config.Version {
		fmt.Printf("json-flatten version %s\n", version)
		os.Exit(ExitOk)
	}

	data := app.Read(config.Input)
	flat := Flatten(data)
	fmt.Fprintf(app.Stdout, "%v\n", flat)
}

func main() {
	app := &Application{
		Args: os.Args,
		Exit: os.Exit,

		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}
	app.Run()
}
