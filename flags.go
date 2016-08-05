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

func parseCommandLine() Config {
	config := Config{}

	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.BoolVar(&config.Help, "h", false, "")
	fs.BoolVar(&config.Version, "v", false, "")

	err := fs.Parse(os.Args[1:])
	if err != nil || config.Help {
		printUsage()
		os.Exit(2)
	}

	if fs.NArg() > 1 {
		printUsage()
		os.Exit(2)
	}

	if fs.NArg() == 1 {
		input := fs.Arg(0)
		config.Input = &input
	}

	return config
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [-v] [-h] [file]\n", os.Args[0])
}
