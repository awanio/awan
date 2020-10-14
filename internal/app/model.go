package app

import (
	"time"

	"github.com/awanio/awan/internal/team"
	"github.com/awanio/awan/internal/user"
	"github.com/awanio/awan/pkg/database"
	"github.com/gofrs/uuid"
)

// Apps [...]
type Apps struct {
	database.BaseModel
	Name           string     `gorm:"unique;column:name;type:varchar(20);not null"`
	CreatedDate    time.Time  `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate    time.Time  `gorm:"column:updated_date;type:datetime;not null"`
	UserID         uuid.UUID  `gorm:"index:user_id;column:user_id;type:uuid;not null"`
	Users          user.Users `gorm:"association_foreignkey:user_id;foreignkey:id"`
	OrganizationID uuid.UUID  `gorm:"index:organization_id;column:organization_id;type:uuid;not null"`
	Organization   team.Teams `gorm:"association_foreignkey:organization_id;foreignkey:id"`
	Type           string     `gorm:"column:type;type:varchar(50)"`
	Zone           string     `gorm:"column:zone;type:varchar(50)"`
	HostServer     string     `gorm:"column:host_server;type:varchar(50)"`
	GitRepo        string     `gorm:"column:git_repo;type:varchar(100)"`
	UUID           string     `gorm:"column:uuid;type:char(36);not null"`
	LastDeployment time.Time  `gorm:"column:last_deployment;type:datetime;not null"`
	LastDeployer   int        `gorm:"index:last_deployer;column:last_deployer;type:int(11)"`
	GitBranch      string     `gorm:"column:git_branch;type:varchar(255)"`
	Replicas       int        `gorm:"column:replicas;type:int(11);not null"`
	InstanceTypeID uuid.UUID  `gorm:"index:instance_type_id;column:instance_type_id;type:uuid;not null"`
	Instances      Instances  `gorm:"association_foreignkey:instance_type_id;foreignkey:id"`
	DeletedDate    time.Time  `gorm:"column:deleted_date;type:datetime"`
	LocalDomain    string     `gorm:"column:local_domain;type:varchar(45)"`
	PublicDomain   string     `gorm:"column:public_domain;type:varchar(45)"`
	Path           string     `gorm:"column:path;type:varchar(45)"`
	Region         string     `gorm:"column:region;type:varchar(255)"`
}

// Metadata [...]
type Metadata struct {
	database.BaseModel
	SourceID       uuid.UUID `gorm:"index:source_id;index:source_id_2;column:source_id;type:uuid;not null"`
	OrganizationID uuid.UUID `gorm:"index:organization_id;column:organization_id;type:uuid"`
	UserID         int       `gorm:"index:user_id;column:user_id;type:int(11)"`
	Type           string    `gorm:"index:source_id_2;column:type;type:enum('app','resource','user','organization');not null"`
	Key            string    `gorm:"index:source_id_2;column:key;type:varchar(255);not null"`
	Value          string    `gorm:"column:value;type:varchar(255);not null"`
	CreatedDate    time.Time `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate    time.Time `gorm:"column:updated_date;type:datetime;not null"`
	DeletedDate    time.Time `gorm:"index:source_id_2;column:deleted_date;type:datetime"`
}

// Instances [...]
type Instances struct {
	database.BaseModel
	CreatedDate  time.Time `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate  time.Time `gorm:"column:updated_date;type:datetime;not null"`
	DeletedDate  time.Time `gorm:"column:deleted_date;type:datetime"`
	InstanceType string    `gorm:"column:instance_type;type:varchar(20);not null"`
	Availability string    `gorm:"column:availability;type:enum('availabile','unavailable');not null"`
	VcpuCore     float32   `gorm:"column:vcpu_core;type:float;not null"`
	RAMMb        int       `gorm:"column:ram_mb;type:int(11);not null"`
	StorageGb    int       `gorm:"column:storage_gb;type:int(11);not null"`
}
