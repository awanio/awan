package model

import (
	"github.com/jinzhu/gorm"
)

// Service data of struct
type Service struct {
	gorm.Model
	ID int `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
}
