package main

import (
	"os"

	"./app"
)

func main() {
	app := app.Get()
	app.Run(os.Args)
}
