package model

import (
	"github.com/awanio/awan/pkg/database"

	"github.com/gofrs/uuid"
)

// Team data of struct
type Team struct {
	database.BaseModel
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
}
