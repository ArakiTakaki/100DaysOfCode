package commands

import (
	"errors"
	"fmt"
	"html/template"
	"os"

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
	if len(args) < 3 {
		return errors.New("値が不正です。")
	}

	table := args[0]
	// record := args[1:]

	// テンプレートファイルの選択
	tmp, err := template.ParseFiles("./template/foundation.go.tmpl")

	// ディレクトリの生成
	os.MkdirAll("./model", os.ModeDir)
	os.MkdirAll("./db", os.ModeDir)

	// ファイルの作成
	file, err := os.Create("./model/" + table + ".go")
	if err != nil {
		fmt.Println("ディレクトリの作成に失敗しました。")
		panic(err)
	}

	// ファイルの書き込み
	err = tmp.Execute(file, nil)

	if err != nil {
		fmt.Println("ファイルの書き込みに失敗しました。")
		panic(err)
	}

	return nil
}
