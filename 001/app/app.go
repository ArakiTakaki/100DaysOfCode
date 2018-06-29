package app

import "github.com/urfave/cli"
import "github.com/ArakiTakaki/100DaysOfCode/001/config"
import "github.com/ArakiTakaki/100DaysOfCode/001/commands"

func Get() *cli.App {
	app := cli.NewApp()
	config.App(app)
	app.Commands = append(app.Commands, commands.Foundation())
	return app
}
