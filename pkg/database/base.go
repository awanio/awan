package database

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// BaseModel contains common columns for all tables.
type BaseModel struct {
	gorm.Model
	ID uuid.UUID `gorm:"column:id;type:uuid;primary_key;not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *BaseModel) BeforeCreate(db *gorm.DB) error {

	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	db.Statement.SetColumn("id", uuid)

	return nil
}
