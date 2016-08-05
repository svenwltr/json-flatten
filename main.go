package main

import (
	"fmt"
	"os"
)

var (
	version = "unknown" // set by makefile
)

func AssertNoError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	config := parseCommandLine()

	if config.Version {
		fmt.Printf("json-flatten version %s\n", version)
		os.Exit(0)
	}

	data := read(config.Input)
	flat := Flatten(data)
	flat.PrintTo(os.Stdout)
}
