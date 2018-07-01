package foundation

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"strings"
)

// FuncMap funcmap
type FuncMap map[string]interface{}

// Create create
func Create(args []string) {
	table := args[0]
	record := args[1:]
	swt := IsExist(table)
	if swt {
		isPanic(errors.New("CreateError"), "既にテーブルが存在しています。")
	}
	createDir()
	createFile(table, record...)
	// ディレクトリの生成
	createMigrate()
}

func createMigrate() {
	file, err := os.OpenFile("./db/history.dat", os.O_RDONLY, 0)
	isPanic(err, "/db/history.datフォルダが読み込めませんでした")
	dat := FileRead(file)
	migrationMessage()
	var work string
	for _, v := range strings.Split(dat, "\n") {
		work += "db.AutoMigrate(&models." + v + "{})"
	}
	tmp, err := template.ParseFiles("./template/migration.go.tmpl")
	isPanic(err, "テンプレートが読み込めませんでした。")

	var builtins = FuncMap{
		"ModelName":  "migration",
		"Attributes": work,
	}

	file, err = os.Create("./db/migration.go")
	isPanic(err, "migrationが作成できませんでした。")
	err = tmp.Execute(file, builtins)
	isPanic(err, "ファイルへ書き込めませんでした。")
	file.Close()
}

// Update update
func Update(args []string) {
	table := args[0]
	record := args[1:]
	swt := IsExist(table)
	if !swt {
		isPanic(errors.New("UpdateError"), "テーブルが見当たりません")
	}
	createDir()
	createFile(table, record...)

}

func createDir() {
	os.MkdirAll("./models", os.ModeDir)
	os.MkdirAll("./db", os.ModeDir)
}

func createFile(table string, record ...string) {
	var attributes string
	tmp, err := template.ParseFiles("./template/foundation.go.tmpl")
	isPanic(err, "テンプレートが読み込めませんでした。")

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
	file, err := os.Create("./models/" + table + ".go")
	isPanic(err, "モデルが作成できませんでした。")
	err = tmp.Execute(file, builtins)
	isPanic(err, "ファイルへ書き込めませんでした。")
	file.Close()
}

// craeteAttributes Table:Attribute:primary_key
func createAttributes(data string) (string, error) {
	work := strings.Split(data, ":")

	if work[1] == "datetime" || work[1] == "timestamp" {
		work[1] = "time.Time"
	}

	data = "\t" + work[0] + " " + work[1]
	if len(work) > 2 {
		data += setProps(work[2:]...)
	}

	data += "\n"
	return data, nil
}

// craeteAttributes Table:Attribute:primary_key
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
