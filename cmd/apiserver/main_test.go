package main

import (
	"fmt"
	"testing"

	"github.com/awanio/awan/internal/user"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/kataras/iris/v12/httptest"
)

type payload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestNewApp(t *testing.T) {

	app := newApp()
	e := httptest.New(t, app)

	factoryUser()

	me, _ := UserRepository.Get()
	strToken, _ := UserRepository.CreateToken(me)

	var (
		expectedPayload = user.Login{
			Username: "Durgan5843",
			Password: "*q3N%+lZnu",
		}

		expectedLogin = map[string]interface{}{
			"code": 200,
			"data": map[string]interface{}{
				"user": map[string]interface{}{
					"username": expectedPayload.Username,
					"status":   "active",
				},
			},
		}

		jwtToken = fmt.Sprintf("Bearer %s", strToken)
	)

	e.GET("/").Expect().Status(httptest.StatusOK)
	e.POST("/api/signin").WithJSON(expectedPayload).Expect().Status(httptest.StatusOK).JSON().Object().ContainsMap(expectedLogin)
	e.GET("/api/users").WithHeader("Authorization", jwtToken).Expect().Status(httptest.StatusOK)
}

// create test user
func factoryUser() (user.Users, error) {

	gofakeit.Seed(0)

	newUser := user.Input{
		Username: gofakeit.Username(),
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, false, 10),
	}

	createdUser, status, err := UserRepository.Create(newUser)

	if err != nil {
		println("Error: ", err.Error())
		return createdUser, err
	}

	if status {
		println("user ID: ", createdUser.ID.String())
		return createdUser, nil
	}

	return createdUser, err

}
