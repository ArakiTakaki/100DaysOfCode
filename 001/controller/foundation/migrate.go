package foundation

import (
	"github.com/ArakiTakaki/100DaysOfCode/001/controller/foundation/databases"
	"github.com/ArakiTakaki/100DaysOfCode/001/db"
	// MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Migration データベースの作成を行う。
func Migration() {
	migrationMessage()
}

func migrationMessage() {
	mysql := databases.UseMySQL()
	mysql.SetDBName("golem")
	mysql.SetUser("root")
	mysql.SetPassword("")
	mysql.SetAddress("localhost")
	mysql.SetCharset("utf8")
	mysql.SetParseTime("true")
	dbacc, err := mysql.GetConn()
	dbacc.LogMode(true)
	isPanic(err, "データベースへ接続できませんでした")
	defer dbacc.Close()
	db.Migration(dbacc)

}
