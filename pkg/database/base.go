package database

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// BaseModel contains common columns for all tables.
type BaseModel struct {
	gorm.Model
	ID        uuid.UUID `gorm:"column:id;type:uuid;primary_key;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null"`
	DeletedAt gorm.DeletedAt
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
