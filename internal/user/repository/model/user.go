package model

import (
	"github.com/jinzhu/gorm"
)

// User data of struct
type User struct {
	gorm.Model
	ID int `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
}
