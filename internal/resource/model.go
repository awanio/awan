package resource

import (
	"time"

	"github.com/awanio/awan/internal/app"
	"github.com/awanio/awan/internal/team"
	"github.com/awanio/awan/internal/user"
	"github.com/awanio/awan/pkg/database"
	"github.com/gofrs/uuid"
)

// Resouce data of struct
type Resource struct {
	database.BaseModel
	Name           string        `gorm:"column:name;type:varchar(100);not null"`
	CreatedDate    time.Time     `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate    time.Time     `gorm:"column:updated_date;type:datetime;not null"`
	Type           string        `gorm:"column:type;type:enum('app','addon');not null"`
	CatalogueID    uuid.UUID     `gorm:"index:catalogue_id;column:catalogue_id;type:uuid;not null"`
	Catalogue      Catalogue     `gorm:"association_foreignkey:catalogue_id;foreignkey:id"`
	UserID         uuid.UUID     `gorm:"index:user_id;column:user_id;type:uuid;not null"`
	Users          user.Users    `gorm:"association_foreignkey:user_id;foreignkey:id"`
	OrganizationID int           `gorm:"index:organization_id;column:organization_id;type:int(11)"`
	Team           team.Teams    `gorm:"association_foreignkey:organization_id;foreignkey:id"`
	Domain         string        `gorm:"column:domain;type:text"`
	InstanceTypeID uuid.UUID     `gorm:"index:instance_type_id;column:instance_type_id;type:uuid;not null"`
	Instances      app.Instances `gorm:"association_foreignkey:instance_type_id;foreignkey:id"`
	DeletedDate    time.Time     `gorm:"column:deleted_date;type:datetime"`
	Replicas       int           `gorm:"column:replicas;type:int(11);not null"`
	Region         string        `gorm:"column:region;type:varchar(255)"`
}

// Catalogue [...]
type Catalogue struct {
	database.BaseModel
	Slug         string    `gorm:"unique;index:slug;column:slug;type:varchar(50);not null"`
	Name         string    `gorm:"column:name;type:varchar(20);not null"`
	Image        string    `gorm:"column:image;type:varchar(100);not null"`
	CreatedDate  time.Time `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate  time.Time `gorm:"column:updated_date;type:datetime;not null"`
	Type         string    `gorm:"column:type;type:enum('app','addon');not null"`
	Port         int       `gorm:"column:port;type:int(11);not null"`
	Logo         string    `gorm:"column:logo;type:longtext"`
	MountPath    string    `gorm:"column:mount_path;type:varchar(200)"`
	PublicAccess bool      `gorm:"column:public_access;type:tinyint(1);not null"`
}
