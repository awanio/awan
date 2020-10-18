package user

// Controller users
type Controller struct{}

// Get method
func (m *Controller) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}
