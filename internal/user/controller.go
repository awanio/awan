package user

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Controller users
type Controller struct {
	Repo RepositoryUsers
}

// Get method
func (m *Controller) Get(ctx iris.Context) {

	me, err := m.Repo.Get()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// println(me.Name)

	if err != nil {
		ctx.JSON(iris.Map{
			"code":  http.StatusBadRequest,
			"error": err.Error,
		})
		return
	}

	ctx.JSON(
		iris.Map{
			"code": http.StatusOK,
			"data": Input{
				Username: me.Username,
				Name:     me.Name,
				Email:    "k4ndar@yahoo.com",
				Password: me.VerificationCode,
			},
		})
}
