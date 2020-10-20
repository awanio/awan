package user

import "gorm.io/gorm"

// Signup controller
type Signup struct {
	DB *gorm.DB
}

// Get method
func (m *Signup) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
