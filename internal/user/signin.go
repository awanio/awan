package user

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// Signin controller
type Signin struct {
	Repo RepositoryUsers
}

// Get method
func (m *Signin) Get() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}

// Post method
func (m *Signin) Post(ctx iris.Context) {

	var login Login

	if err := ctx.ReadJSON(&login); err != nil {
		println(err.Error())
		ctx.StopWithError(iris.StatusBadRequest, err)
	}

	me, token, status, er := m.Repo.Authenticate(login.Username, login.Password)

	if er != nil {
		ctx.JSON(iris.Map{
			"code":  http.StatusBadRequest,
			"error": er.Error,
		})
		return
	}

	if !status {
		ctx.JSON(iris.Map{
			"code":  http.StatusUnauthorized,
			"error": "login fail",
		})
		return
	}

	ctx.JSON(
		iris.Map{
			"code": http.StatusOK,
			"data": map[string]interface{}{
				"user": map[string]interface{}{
					"id":       me.ID,
					"username": me.Username,
					"name":     me.Name,
					"status":   me.Status,
				},
				"token": token,
			},
		})
}
