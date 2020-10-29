package user

import (
	"github.com/awanio/awan/pkg/helper"
	"github.com/iris-contrib/middleware/jwt"
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

// GetByUsername ...
func (m *RepositoryUser) GetByUsername(username string) (Users, error) {

	var existingUser Users

	result := m.DB.Where("username = ?", username).First(&existingUser)

	if result.Error != nil {
		return existingUser, result.Error
	}

	return existingUser, nil
}

// Authenticate method
func (m *RepositoryUser) Authenticate(username string, password string) (Users, string, bool, error) {

	existingUser, err := m.GetByUsername(username)

	if err != nil {
		return existingUser, "", false, err
	}

	var credential Credentials

	result := m.DB.Where("user_id = ? and type = ?", existingUser.ID, "password").First(&credential)

	if result.Error != nil {
		return existingUser, "", false, result.Error
	}

	jwtToken, _ := m.CreateToken(existingUser)

	return existingUser, jwtToken, true, nil
}

// CreateToken ...
func (m *RepositoryUser) CreateToken(user Users) (string, error) {

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte("My Secret"))

}
