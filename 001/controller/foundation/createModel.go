package foundation

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/ArakiTakaki/100DaysOfCode/001/app/util"
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
		work += "db.AutoMigrate(&models." + v + "{})\n"
	}

	fmt.Println(work)
	f := util.GetTextGenerater()
	f.SetFilePath("./db/migration.go")
	f.SetTmplate("./template/migration.go.tmpl")
	f.SetTitle("migration.go")
	f.SetContent(work)
	err = f.Excute()
	isPanic(err, "ファイルが生成できませんでした:migration")
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
