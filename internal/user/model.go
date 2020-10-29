package user

import (
	"time"

	"github.com/awanio/awan/pkg/database"
	"github.com/gofrs/uuid"
)

// Users data of struct
type Users struct {
	database.BaseModel
	Username                    string    `gorm:"unique;column:username;type:varchar(20);not null"`
	CreatedDate                 time.Time `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate                 time.Time `gorm:"column:updated_date;type:datetime;not null"`
	DeletedDate                 time.Time `gorm:"column:deleted_date;type:datetime;not null"`
	Name                        string    `gorm:"column:name;type:varchar(30)"`
	Status                      string    `gorm:"column:status;type:varchar(20);not null"`
	LastLogin                   time.Time `gorm:"column:last_login;type:datetime"`
	VerificationCode            string    `gorm:"unique;column:verification_code;type:varchar(8)"`
	ForgotPasswordCode          string    `gorm:"unique;column:forgot_password_code;type:varchar(8)"`
	ForgotPasswordCodeValidTime time.Time `gorm:"column:forgot_password_code_valid_time;type:datetime"`
}

// Input struct
type Input struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Credentials [...]
type Credentials struct {
	database.BaseModel
	UserID           uuid.UUID `gorm:"index:user_id;column:user_id;type:uuid;not null"`
	Users            Users     `gorm:"association_foreignkey:user_id;foreignkey:id"`
	Type             string    `gorm:"column:type;type:varchar(20);not null"` // 'password','email','phone','sshkey','vendor','fingerprint'
	UserKey          string    `gorm:"column:user_key;type:text"`
	UserValue        string    `gorm:"column:user_value;type:text"`
	Status           string    `gorm:"column:status;type:varchar(20);not null"` // enum('active','inactive','unverified','deleted')
	CreatedDate      time.Time `gorm:"column:created_date;type:datetime;not null"`
	UpdatedDate      time.Time `gorm:"column:updated_date;type:datetime;not null"`
	DeletedDate      time.Time `gorm:"column:deleted_date;type:datetime;not null"`
	LastAccessedDate time.Time `gorm:"column:last_accessed_date;type:datetime;not null"`
	Desc             string    `gorm:"column:desc;type:text"`
}

// Login ...
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
