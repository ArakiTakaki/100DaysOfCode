package models

import (
	// migrationエラーの回避
	_ "time"
	"github.com/jinzhu/gorm"
)

type Sasisusseso struct {
	gorm.Model
		Name string

}