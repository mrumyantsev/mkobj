package main

import (
	"github.com/mrumyantsev/mkobj/internal/app/mkdir"
	"github.com/mrumyantsev/mkobj/internal/output"
)

func main() {
	app, err := mkdir.New()
	if err != nil {
		output.Fatal(err)
	}

	if err = app.Run(); err != nil {
		output.Fatal(err)
	}
}
