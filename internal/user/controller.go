package user

import (
	"gorm.io/gorm"
)

// Controller users
type Controller struct{}

// Get method
func (m *Controller) Get(db *gorm.DB) interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
