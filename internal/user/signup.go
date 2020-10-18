package user

// Signup serves the "/", "/ping" and "/hello".
type Signup struct{}

// Get method
func (m *Signup) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
