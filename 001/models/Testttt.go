package models

import (
	// migrationエラーの回避
	_ "time"
	"github.com/jinzhu/gorm"
)

type Testttt struct {
	gorm.Model
		Name string

}