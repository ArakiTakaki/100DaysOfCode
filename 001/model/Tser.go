package models

import (
	"time"
	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
// TODO ドライバの指定の仕方は考える必要がある。

type Tser struct {
	gorm.Model
	
	TestCel int
	TestCel2 string
	created_at time.Time
}