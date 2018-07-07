package models

import (
	// migrationエラーの回避
	_ "time"
	"github.com/jinzhu/gorm"
)

type Testttttt6 struct {
	gorm.Model
		Name string

}