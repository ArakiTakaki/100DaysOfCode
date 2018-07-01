package db

import (
	"github.com/ArakiTakaki/100DaysOfCode/001/models"
	"github.com/jinzhu/gorm"
	// MySQL
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Test{})

}
