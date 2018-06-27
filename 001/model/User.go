package models

import (
	"time"
	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// TODO ドライバの指定の仕方は考える必要がある。

type User struct {
	gorm.Model
		Id string
	Pass string
	Name string `gorm:&#34;not null&#34;`
	BitrhDay time.Time `gorm:&#34;not_null1&#34;`

}