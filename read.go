package main

import (
	"encoding/json"
	"os"
)

func read(filename *string) interface{} {
	input := os.Stdin

	if filename != nil {
		var err error
		input, err = os.Open(*filename)
		AssertNoError(err)
	}

	var data interface{}

	decoder := json.NewDecoder(input)
	err := decoder.Decode(&data)
	AssertNoError(err)

	return data
}
