package user

import (
	"net/http"

	"github.com/awanio/awan/pkg/helper"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

// Signup controller
type Signup struct {
	DB *gorm.DB
}

// Get method
func (m *Signup) Get() interface{} {

	return map[string]string{"message": "Hello Signup Controller!"}
}

// Post method
func (m *Signup) Post(ctx iris.Context) {

	verificationCode, _ := helper.GenerateRandomString(8)
	forgotPasswordCode, _ := helper.GenerateRandomString(8)

	err := m.DB.Create(&Users{
		Username:           "test_" + verificationCode,
		Name:               "Iskandar Soesman",
		Status:             "active",
		VerificationCode:   verificationCode,
		ForgotPasswordCode: forgotPasswordCode,
	})

	if err != nil {
		println(err.Error)

	}

	var createdUser Users

	if err := m.DB.First(&createdUser, "username = ?", "k4ndar"); err == nil {
		println(err.Error)
		// app.Logger().Fatalf("created one record failed: %s", err.Error)

		ctx.JSON(iris.Map{
			"code":  http.StatusBadRequest,
			"error": err.Error,
		})

	}

	ctx.JSON(
		iris.Map{
			"code": http.StatusOK,
			"data": Input{
				Username: createdUser.Username,
				Name:     createdUser.Name,
				Email:    "k4ndar@yahoo.com",
				Password: createdUser.VerificationCode,
			},
		})
}
