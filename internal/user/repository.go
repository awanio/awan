package user

import (
	"fmt"

	"github.com/awanio/awan/internal/env"
	"github.com/awanio/awan/pkg/helper"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gofrs/uuid"
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
	Authenticate(string, string) (Users, string, bool, error)
	CreateAdmin() (map[string]string, bool, error)
	CreateToken(user Users) (string, error)
	Create(newUser Input) (Users, bool, error)
	Delete(uuid uuid.UUID)
}

// NewRepository ...
func NewRepository(db *gorm.DB) *RepositoryUser {

	return &RepositoryUser{DB: db, Status: true}
}

// Get ...
func (m *RepositoryUser) Get() (Users, error) {

	var existingUser Users

	result := m.DB.First(&existingUser)

	if result.Error != nil {
		println("error")
		println(result.Error)
		return existingUser, result.Error
	}

	return existingUser, nil
}

// Delete a user
func (m *RepositoryUser) Delete(uuid uuid.UUID) {

	tx := m.DB.Begin()
	tx.Delete(&Users{}, uuid)
	tx.Where("user_id = ?", uuid).Delete(&Credentials{}, uuid)
	tx.Commit()
}

// Create for a user
func (m *RepositoryUser) Create(newUser Input) (Users, bool, error) {

	// does this username exists?
	existingUser, err := m.GetByUsername(newUser.Username)

	if err == nil {
		// yes, username with this value exists
		return existingUser, false, nil
	}

	// does this email exists?
	existingEmail, err := m.GetByEmail(newUser.Email)

	if err == nil {
		// yes, email with this value exists
		return existingEmail, false, nil
	}

	verificationCode, _ := helper.GenerateRandomString(8)
	forgotPasswordCode, _ := helper.GenerateRandomString(8)

	var user = Users{
		Username:           newUser.Username,
		Name:               newUser.Name,
		Status:             "active",
		VerificationCode:   verificationCode,
		ForgotPasswordCode: forgotPasswordCode,
	}

	tx := m.DB.Begin()
	create := tx.Create(&user)

	if create.Error != nil {
		tx.Rollback()
		return existingUser, false, create.Error

	}

	passwd := []byte(newUser.Password)
	hashedPasswd, erro := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost)

	if erro != nil {
		tx.Rollback()
		return user, false, erro
	}

	result := tx.Create(&Credentials{
		UserID:  user.ID,
		Type:    "password",
		UserKey: string(hashedPasswd),
		Status:  "active",
	})

	if result.Error != nil {
		tx.Rollback()
		return user, false, result.Error

	}

	tx.Commit()

	return user, true, nil
}

// CreateAdmin for the first time only
func (m *RepositoryUser) CreateAdmin() (map[string]string, bool, error) {

	gofakeit.Seed(0)

	adminPassword := gofakeit.Password(true, true, true, true, false, 10)
	adminUsername := gofakeit.Username()

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

	var admin = Input{
		Username: adminUsername,
		Name:     "Admin",
		Email:    "active",
		Password: adminPassword,
	}

	_, status, err := m.Create(admin)

	if !status {
		return resp, false, err
	}

	if err != nil {
		return resp, false, err
	}

	return resp, true, nil
}

// GetByUsername ...
func (m *RepositoryUser) GetByUsername(username string) (Users, error) {

	return m.GetBy("username", username)
}

// GetByEmail ...
func (m *RepositoryUser) GetByEmail(email string) (Users, error) {

	return m.GetBy("email", email)
}

// GetByID ...
func (m *RepositoryUser) GetByID(uuid uuid.UUID) (Users, error) {

	return m.GetBy("id", uuid.String())
}

// GetBy ...
func (m *RepositoryUser) GetBy(column, value string) (Users, error) {

	var existingUser Users

	where := fmt.Sprintf("%s = ?", column)
	result := m.DB.Where(where, value).First(&existingUser)

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

	if !m.checkPasswordHash(password, credential.UserKey) {
		return existingUser, "", false, nil
	}

	jwtToken, _ := m.CreateToken(existingUser)

	return existingUser, jwtToken, true, nil
}

func (m *RepositoryUser) checkPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateToken ...
func (m *RepositoryUser) CreateToken(user Users) (string, error) {

	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(env.JWTSecret))

}
