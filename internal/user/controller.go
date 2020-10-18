package user

import (
	"github.com/kataras/iris"
)

// Controller serves the "/", "/ping" and "/hello".
type Controller struct{}

// Get method
func (m *Controller) Get(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
}
