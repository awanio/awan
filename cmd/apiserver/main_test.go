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
	)

	e.GET("/").Expect().Status(httptest.StatusOK)
	e.POST("/api/signin").WithJSON(expectedPayload).Expect().Status(httptest.StatusOK)

	e.GET("/api/users").Expect().Status(httptest.StatusOK)
	// e.GET("/api/apps").Expect().Status(httptest.StatusOK)
	// e.GET("/api/resources").Expect().Status(httptest.StatusOK)
	// e.GET("/api/teams").Expect().Status(httptest.StatusOK)

	// json example
	// expectedErr := map[string]interface{}{
	// 	"app":     app.AppName,
	// 	"status":  httptest.StatusNotFound,
	// 	"message": "",
	// }
	// e.GET("/anotfoundwithjson").WithQuery("json", nil).
	// 	Expect().Status(httptest.StatusNotFound).JSON().Equal(expectedErr)
}
