package models

import (
	// migrationエラーの回避
	_ "time"
	"github.com/jinzhu/gorm"
)

type Testttttt struct {
	gorm.Model
		Name string

}