package team

import (
	"time"

	"github.com/awanio/awan/internal/user"
	"github.com/awanio/awan/pkg/database"
	"github.com/gofrs/uuid"
)

// Teams data of struct
type Teams struct {
	database.BaseModel
	Name        string    `gorm:"unique;column:name;type:varchar(20);not null"`
	Slug        string    `gorm:"column:slug;type:varchar(50);not null"`
	CreatedDate time.Time `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate time.Time `gorm:"column:updated_date;type:datetime;not null"`
}

// TeamMembers [...]
type TeamMembers struct {
	database.BaseModel
	CreatedDate time.Time  `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate time.Time  `gorm:"column:updated_date;type:datetime;not null"`
	UserID      uuid.UUID  `gorm:"index:user_id;column:user_id;type:uuid;not null"`
	Users       user.Users `gorm:"association_foreignkey:user_id;foreignkey:id"`
	TeamID      uuid.UUID  `gorm:"unique;column:team_id;type:uuid;not null"`
	Team        Teams      `gorm:"association_foreignkey:team_id;foreignkey:id"`
	UserLevel   int        `gorm:"column:user_level;type:int(11);not null"`
}
