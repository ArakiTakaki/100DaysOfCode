package databases

// 各種SQLのドライバを定義
const (
	mysqlDriver    = "github.com/jinzhu/gorm/dialects/mysql"
	PostgresDriver = "github.com/jinzhu/gorm/dialects/postgres"
	MSSqlDriver    = "github.com/jinzhu/gorm/dialects/mssql"
	SQLiteDriver   = "github.com/jinzhu/gorm/dialects/sqlite"
)
