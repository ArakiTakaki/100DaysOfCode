package databases

import "github.com/jinzhu/gorm"

type SQLite struct {
	Dir string
}

// UseSQLite SQLite3を使用します。
// param dir DBの配置されているディレクトリを記載してください。
func UseSQLite(dir string) SQLite {
	var d SQLite
	d.Dir = dir
	return d
}

// GetConn コネクションするときに用います。
func (d SQLite) GetConn() (*gorm.DB, error) {
	return gorm.Open("sqlite3", d.Dir)
}
