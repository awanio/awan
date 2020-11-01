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

	me, strToken, credential, _ := factoryUser()

	var (
		expectedPayload = user.Login{
			Username: credential["username"],
			Password: credential["password"],
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

	UserRepository.Delete(me.ID)
}

// create test user
func factoryUser() (user.Users, string, map[string]string, error) {

	gofakeit.Seed(0)

	username := gofakeit.Username()
	password := gofakeit.Password(true, true, true, true, false, 10)
	credential := map[string]string{
		"username": username,
		"password": password,
	}

	newUser := user.Input{
		Username: username,
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: password,
	}

	createdUser, status, err := UserRepository.Create(newUser)

	if err != nil {
		println("Error: ", err.Error())
		return createdUser, "", credential, err
	}

	strToken, _ := UserRepository.CreateToken(createdUser)

	if status {
		return createdUser, strToken, credential, nil
	}

	return createdUser, "", credential, err

}
