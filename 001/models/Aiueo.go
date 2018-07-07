package models

import (
	// migrationエラーの回避
	_ "time"
	"github.com/jinzhu/gorm"
)

type Aiueo struct {
	gorm.Model
		Name string

}