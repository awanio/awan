package user

import (
	"github.com/kataras/iris/v12"
)

// Controller serves the "/", "/ping" and "/hello".
type Controller struct{}

// Get method
func (m *Controller) Get(ctx iris.Context) {

	ctx.JSON(Input{
		Username: "Makis",
		Name:     "Makis",
		Email:    "Makis",
	})
}
