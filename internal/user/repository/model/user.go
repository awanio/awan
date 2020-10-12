package model

import (
	"github.com/awanio/awan/pkg/database"
	"github.com/gofrs/uuid"
)

// User data of struct
type User struct {
	database.BaseModel
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string
	Username string
	Status   string
}
