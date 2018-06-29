package foundation

import (
	"os"
)

// Migration データベースの作成を行う。
func Migration() {
	file, err := os.OpenFile("./db/history.dat", os.O_RDONLY, 0)
	isPanic(err, "マイグレーションが行えませんでした")
	FileRead(file)
}

func migrationMessage() {
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
