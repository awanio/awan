package user

import (
	"gorm.io/gorm"
)

// Controller users
type Controller struct {
	DB *gorm.DB
}

// Get method
func (m *Controller) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
