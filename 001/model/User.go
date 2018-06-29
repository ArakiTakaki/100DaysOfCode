package models

import (
	"time"

	"github.com/jinzhu/gorm"
	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO ドライバの指定の仕方は考える必要がある。

// User user
type User struct {
	gorm.Model
	ID       string
	Pass     string
	Name     string
	BitrhDay time.Time
}
