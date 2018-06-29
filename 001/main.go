package main

import (
	"os"

	"github.com/ArakiTakaki/100DaysOfCode/001/app"
)

func main() {
	app := app.Get()
	app.Run(os.Args)
}
