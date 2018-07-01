package databases

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

//params : https://github.com/go-sql-driver/mysql#parameters

// MySQL MySQLへの接続するためのパラメータ郡の策定を行う
type MySQL struct {
	User     string
	Password string
	Address  string
	DBName   string
	Params   []string
}

// UseMySQL mysqlを使用します。
func UseMySQL() MySQL {
	return MySQL{}
}

// SetUser user
func (d *MySQL) SetUser(userName string) {
	d.User = userName
}

// SetPassword pass
func (d *MySQL) SetPassword(password string) {
	d.Password = password
}

// SetAddress adr
func (d *MySQL) SetAddress(data string) {
	if data != "localhost" {
		d.Address = "tcp(" + data + ":3306)"
	}
	d.Address = ""
}

// SetDBName dbname
func (d *MySQL) SetDBName(dbname string) {
	d.DBName = dbname
}

func (d *MySQL) parse() string {
	var work string
	work += d.User + ":" + d.Password + "@"
	work += d.Address
	work += "/" + d.DBName + "?"
	if len(d.Params) != 0 {
		work += strings.Join(d.Params, "&")
	}
	fmt.Println(work)
	return work
}

// GetConn コネクションを確立
func (d *MySQL) GetConn() (*gorm.DB, error) {
	return gorm.Open("mysql", d.parse())
}

// SetCharset utf8 sjis latin1 utf8mb4,utf8 等など
func (d *MySQL) SetCharset(data string) {
	d.Params = append(d.Params, "charset="+data)
}

// SetParseTime true false
func (d *MySQL) SetParseTime(data string) {
	d.Params = append(d.Params, "parseTime="+data)
}

// SetLoc Local
func (d *MySQL) SetLoc(data string) {
	d.Params = append(d.Params, "loc="+data)
}

// SetCollation 称号順序 default:utf8_general_ci
func (d *MySQL) SetCollation(data string) {
	d.Params = append(d.Params, "collation="+data)
}
