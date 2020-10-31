package main

import (
	"fmt"
	"testing"

	"github.com/awanio/awan/internal/user"
	"github.com/kataras/iris/v12/httptest"
)

type payload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestNewApp(t *testing.T) {

	app := newApp()
	e := httptest.New(t, app)

	me, _ := UserRepository.Get()
	strToken, _ := UserRepository.CreateToken(me)

	var (
		expectedPayload = user.Login{
			Username: "t6oJq",
			Password: "dR7AniI2",
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
