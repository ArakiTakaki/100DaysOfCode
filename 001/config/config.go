package config

import "github.com/urfave/cli"

// Config アプリのコンフィグ
func App(app *cli.App) {
	app.Name = "boom"
	app.Usage = "make an explosive entrrance"
	app.Version = "0.0.1"
}
