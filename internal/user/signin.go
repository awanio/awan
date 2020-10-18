package user

// Signin controller
type Signin struct{}

// Get method
func (m *Signin) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
