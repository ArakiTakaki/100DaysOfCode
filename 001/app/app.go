package app

import "github.com/urfave/cli"
import "../config"
import "../commands"

func Get() *cli.App {
	app := cli.NewApp()
	config.App(app)
	app.Commands = append(app.Commands, commands.Foundation())
	return app
}
