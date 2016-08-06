package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Help    bool
	Version bool
	Input   *string
}

func (app *Application) GetConfig() Config {
	config := Config{}

	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.BoolVar(&config.Help, "h", false, "")
	fs.BoolVar(&config.Version, "v", false, "")

	err := fs.Parse(app.Args[1:])
	if err != nil || config.Help {
		app.Usage()
	}

	if fs.NArg() > 1 {
		app.Usage()
	}

	if fs.NArg() == 1 {
		input := fs.Arg(0)
		config.Input = &input
	}

	return config
}

func (app *Application) Usage() {
	fmt.Fprintf(app.Stderr, "Usage: %s [-v] [-h] [file]\n", app.Args[0])
	app.Exit(ExitUsage)
}
