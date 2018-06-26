package commands

import (
	"fmt"
	"html/template"
	"os"
	"strings"

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
type FuncMap map[string]interface{}

func action(c *cli.Context) error {
	args := c.Args()

	table := args[0]
	record := args[1:]

	os.MkdirAll("./model", os.ModeDir)
	os.MkdirAll("./db", os.ModeDir)

	createFile(table, record...)
	// ディレクトリの生成

	return nil
}

func createFile(table string, record ...string) {

	var attributes string
	// テンプレートファイルの選択
	// ファイルの作成

	tmp, err := template.ParseFiles("./template/foundation.go.tmpl")
	isPanic(err, "テンプレートが読み込めませんでした。")

	file, err := os.Create("./model/" + table + ".go")
	isPanic(err, "モデルが作成できませんでした。")

	for _, v := range record {
		value, err := createAttributes(v)
		attributes += value
		isPanic(err, "createAttributesでエラーを起こしました。")
	}

	var builtins = FuncMap{
		"ModelName":  table,
		"Attributes": attributes,
	}
	// ファイルの書き込み
	err = tmp.Execute(file, builtins)
	isPanic(err, "ファイルへ書き込めませんでした。")

}

// craeteAttributes Table:Attribute:primary_key
func createAttributes(data string) (string, error) {
	work := strings.Split(data, ":")
	data = "\t" + work[0] + " " + work[1]
	if len(work) > 2 {
		data += setProps(work[2:]...)
	}

	data += "\n"
	return data, nil
}

func setProps(data ...string) string {

	for i, s := range data {
		if s == "not_null" {
			data[i] = "not null"
		}
	}

	var work string
	work += " `gorm:"
	work += "\""
	work += strings.Join(data, ";")
	work += "\""
	work += "`"
	fmt.Println(work)
	return work
}

func isPanic(work error, message string) {
	if work != nil {
		fmt.Println(message)
		panic(work)
	}
}
