package main

import (
	"encoding/json"
	"os"
)

func (app *Application) Read(filename *string) interface{} {
	input := app.Stdin

	if filename != nil {
		var err error
		input, err = os.Open(*filename)
		app.AssertNoError(err)
	}

	var data interface{}

	decoder := json.NewDecoder(input)
	err := decoder.Decode(&data)
	app.AssertNoError(err)

	return data
}
