package model

import (
	"github.com/jinzhu/gorm"
)

// Team data of struct
type Team struct {
	gorm.Model
	ID int `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
}
