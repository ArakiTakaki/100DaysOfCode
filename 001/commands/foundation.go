package commands

import (
	"github.com/ArakiTakaki/100DaysOfCode/001/controller/foundation"
	"github.com/urfave/cli"
)

// Foundation テスト
func Foundation() cli.Command {
	foundation := cli.Command{
		Name:   "foundation",
		Usage:  "DBの生成、編集を行います",
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

	if status == "migration" || status == "m" {
		foundation.Migration()
	}

	return nil
}
