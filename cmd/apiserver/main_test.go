package main

import (
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestNewApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)

	e.GET("/").Expect().Status(httptest.StatusOK)
	e.GET("/api/signin").Expect().Status(httptest.StatusOK)
	e.GET("/api/users").Expect().Status(httptest.StatusOK)
	e.GET("/api/apps").Expect().Status(httptest.StatusOK)
	e.GET("/api/resources").Expect().Status(httptest.StatusOK)
	e.GET("/api/teams").Expect().Status(httptest.StatusOK)
}
