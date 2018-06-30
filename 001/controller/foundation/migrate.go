package foundation

import (
	"github.com/ArakiTakaki/100DaysOfCode/001/controller/foundation/databases"
	"github.com/ArakiTakaki/100DaysOfCode/001/models"
	// MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Migration データベースの作成を行う。
func Migration() {
	//file, err := os.OpenFile("./db/history.dat", os.O_RDONLY, 0)
	//isPanic(err, "マイグレーションが行えませんでした")
	//test := FileRead(file)
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
	db, err := mysql.GetConn()
	isPanic(err, "データベースへ接続できませんでした")
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	defer db.Close()

	db.AutoMigrate(&models.Test{})

	// MySQL
	// // postgres
	// db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	// defer db.Close()
	// host=myhost port=myport user=gorm dbname=gorm password=mypassword

	// // SQLite3
	// db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	// defer db.Close()

	// // MSSql
	// db, err = gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
	// defer db.Close()

}
