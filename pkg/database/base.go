package database

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// BaseModel contains common columns for all tables.
type BaseModel struct {
	gorm.Model
	ID uuid.UUID `gorm:"index:id;column:id;type:uuid;primary_key;not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("id", uuid)
}
