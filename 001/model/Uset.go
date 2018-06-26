package models

import (
	"time"
	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// TODO ドライバの指定の仕方は考える必要がある。

type Uset struct {
	gorm.Model
		Test string `gorm:&#34;not null&#34;`

}