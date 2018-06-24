package commands

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

/**
 * 1 DBName
 */
func Foundation() cli.Command {
	foundation := cli.Command{
		Name:   "foundation",
		Usage:  "DBの生成を行います",
		Action: action,
	}
	return foundation
}

func action(c *cli.Context) error {
	test := c.Args()
	if len(test) < 3 {
		return errors.New("値が不正です。")
	}
	// サンプル
	fmt.Println("type " + test[0] + " struct{\n  gorm.Model\n  " + test[1] + ": " + test[2] + "\n}")
	return nil
}
