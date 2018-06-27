package commands

import (
	"fmt"

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

// FuncMap test
// @TODO 今度フラグ作ろう というか多分newとかそこらへんフラグで管理をするんだろうなぁ
func action(c *cli.Context) error {
	args := c.Args()
	status := args[0]
	params := args[1:]

	if status == "new" {
		stat := foundation.IsExist(params[0])
		if stat {
			foundation.Create(params)
			foundation.AddHistroy(params[0])
		} else {
			fmt.Println("既にテーブルが存在しています、変更したい場合`update`を使用してください。")
		}
	}

	if status == "update" {
		foundation.Create(params)
	}

	if status == "migrate" {
	}

	return nil
}
