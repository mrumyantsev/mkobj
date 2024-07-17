package main

import (
	"github.com/mrumyantsev/mkobj/internal/app"
	"github.com/mrumyantsev/mkobj/internal/constants"
	"github.com/mrumyantsev/mkobj/pkg/clog"
)

func main() {
	a, err := app.New()
	if err != nil {
		clog.Fatal(constants.AppPrintMark + err.Error())
	}

	if err = a.Run(); err != nil {
		clog.Fatal(constants.AppPrintMark + err.Error())
	}
}
