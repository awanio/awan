package user

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Controller users
type Controller struct {
	repo RepositoryUsers
}

// Get method
func (m *Controller) Get(ctx iris.Context) {

	// result, err := m.Users.Get()
	// me, err := m.repo.Get()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// println(me.Name)

	// if err != nil {
	// 	ctx.JSON(iris.Map{
	// 		"code":  http.StatusBadRequest,
	// 		"error": err.Error,
	// 	})
	// }

	ctx.JSON(
		iris.Map{
			"code": http.StatusOK,
			"data": Input{
				// Username: result.Username,
				// Name:     result.Name,
				Email: "k4ndar@yahoo.com",
				// Password: result.VerificationCode,
			},
		})
}
