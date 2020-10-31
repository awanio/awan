package main

import (
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
	)
	jwtToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjU2ZTY3MWNmLWI1MzAtNGRkNC05MzhjLTBkZWNmYjU5MjJkYyIsInVzZXJuYW1lIjoidDZvSnEifQ.k8uILws5SIkNew6O3u1SOxp-QDyRoxv6x0_noySJg84"

	e.GET("/").Expect().Status(httptest.StatusOK)
	resp := e.POST("/api/signin").WithJSON(expectedPayload).Expect().Status(httptest.StatusOK)
	println("json response", resp.JSON().Object().ContainsMap(expectedLogin))

	e.GET("/api/users").WithHeader("Authorization", jwtToken).Expect().Status(httptest.StatusOK)

	// e.GET("/api/apps").Expect().Status(httptest.StatusOK)
	// e.GET("/api/resources").Expect().Status(httptest.StatusOK)
	// e.GET("/api/teams").Expect().Status(httptest.StatusOK)
}
