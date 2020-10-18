package user

// Signin serves the "/", "/ping" and "/hello".
type Signin struct{}

// Get method
func (m *Signin) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
