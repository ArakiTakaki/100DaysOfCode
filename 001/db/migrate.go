package db

import (
	"github.com/jinzhu/gorm"
	// MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Migration(db *gorm.DB) {
}
