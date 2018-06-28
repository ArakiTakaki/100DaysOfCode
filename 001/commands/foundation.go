package commands

import (
	"../controller/foundation"
	"github.com/urfave/cli"
)

// Foundation テスト
func Foundation() cli.Command {
	foundation := cli.Command{
		Name:   "foundation",
		Usage:  "DBの生成を行います",
		Action: action,
	}
	return foundation
}

func action(c *cli.Context) error {
	args := c.Args()
	status := args[0]
	params := args[1:]

	if status == "new" || status == "n" {
		foundation.Create(params)
		foundation.AddHistroy(params[0])
	}
	if status == "update" || status == "u" {
		foundation.Update(params)
	}

	if status == "migrate" {
	}

	return nil
}
