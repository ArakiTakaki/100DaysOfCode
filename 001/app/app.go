package app

import (
	"github.com/ArakiTakaki/100DaysOfCode/001/commands"
	"github.com/ArakiTakaki/100DaysOfCode/001/config"
	"github.com/urfave/cli"
)

func Get() *cli.App {
	app := cli.NewApp()
	config.App(app)
	app.Commands = append(app.Commands, commands.Foundation())
	return app
}
