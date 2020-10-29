package user

import (
	"github.com/awanio/awan/pkg/helper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RepositoryUser ...
type RepositoryUser struct {
	DB     *gorm.DB
	Status bool
}

// RepositoryUsers ...
type RepositoryUsers interface {
	Get() (user Users, err error)
}

// NewRepository ...
func NewRepository(db *gorm.DB) *RepositoryUser {

	return &RepositoryUser{DB: db, Status: true}
}

// Get ...
func (m *RepositoryUser) Get() (Users, error) {

	var existingUser Users

	err := m.DB.First(&existingUser)

	if err != nil {
		return existingUser, err.Error
	}

	return existingUser, nil
}

// CreateAdmin for the first time only
func (m *RepositoryUser) CreateAdmin() (map[string]string, bool, error) {

	verificationCode, _ := helper.GenerateRandomString(8)
	forgotPasswordCode, _ := helper.GenerateRandomString(8)
	adminPassword, _ := helper.GenerateRandomString(8)
	adminUsername, _ := helper.GenerateRandomString(5)

	resp := map[string]string{
		"adminPassword": adminPassword,
		"adminUsername": adminUsername,
	}

	var user Users
	result := m.DB.First(&user)

	// admin already created
	if result.RowsAffected > 0 {
		return resp, false, result.Error
	}

	var admin = Users{
		Username:           adminUsername,
		Name:               "Admin",
		Status:             "active",
		VerificationCode:   verificationCode,
		ForgotPasswordCode: forgotPasswordCode,
	}

	create := m.DB.Create(&admin)

	if create.Error != nil {
		return resp, false, create.Error

	}

	passwd := []byte(adminPassword)
	hashedPasswd, erro := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)

	if erro != nil {
		return resp, false, erro
	}

	result = m.DB.Create(&Credentials{
		UserID:  admin.ID,
		Type:    "password",
		UserKey: string(hashedPasswd),
		Status:  "active",
	})

	if result.Error != nil {
		return resp, false, result.Error

	}

	return resp, true, nil
}
