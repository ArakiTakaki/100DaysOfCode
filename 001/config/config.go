package config

import (
	"github.com/urfave/cli"
)

// App アプリのコンフィグ
func App(app *cli.App) {
	app.Name = "interact"
	app.Usage = "頑張ろう"
	app.Version = "0.0.1"
	app.UsageText = "teststeste"
}
