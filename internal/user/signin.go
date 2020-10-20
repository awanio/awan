package user

import "gorm.io/gorm"

// Signin controller
type Signin struct {
	DB *gorm.DB
}

// Get method
func (m *Signin) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
