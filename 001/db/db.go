package models

import (
	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// NewDBConn (DBの接続を返す)
func newDBConn() *gorm.DB {
	const CONNECT = "root@/shingle_lesson"
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		panic(err)
	}
	return db
}

// Get DBの接続を返す（基本的に
func Get() *gorm.DB {
	db = newDBConn()
	return db
}
