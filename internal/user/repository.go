package user

import "gorm.io/gorm"

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

	// var me Users
	// me.Username = "test_"
	// me.Name = "Iskandar"

	// return me, nil

	var existingUser Users

	err := m.DB.First(&existingUser, "username = ?", "k4ndar")

	if err != nil {
		return existingUser, err.Error
	}

	return existingUser, nil
}
