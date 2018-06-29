package databases

import "github.com/jinzhu/gorm"

type MySQL struct {
	DBName    string
	User      string
	Password  string
	ParseTime string
	Location  string
	Address   string
	params    []string
}

// UseMySQL mysqlを使用します。
func UseMySQL() MySQL {
	return MySQL{}
}

// GetConn コネクションを確立
func (d MySQL) GetConn() (*gorm.DB, error) {
	return gorm.Open("mysql", d.DBName)
}
