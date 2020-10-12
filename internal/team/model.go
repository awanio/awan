package team

import (
	"github.com/awanio/awan/pkg/database"
	"github.com/gofrs/uuid"
)

// Service data of struct
type Service struct {
	database.BaseModel
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
}
