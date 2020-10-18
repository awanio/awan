package user

// Controller serves the "/", "/ping" and "/hello".
type Controller struct{}

// Get method
func (m *Controller) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
