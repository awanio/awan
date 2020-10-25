package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/awanio/awan/internal/db"
	"github.com/awanio/awan/internal/env"
	"github.com/awanio/awan/internal/user"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func init() {
	envFileName := ".env.example"
	env.Load(envFileName)
}

func newApp() *iris.Application {

	app := iris.Default()
	app.Logger().SetLevel("debug")

	connectedDB, sqlDB, err := db.Run()

	if err != nil {
		app.Logger().Fatalf("connect to sqlite3 failed")
		return nil
	}

	iris.RegisterOnInterrupt(func() {
		defer sqlDB.Close()
	})

	err = connectedDB.AutoMigrate(&user.Users{}, &user.Credentials{})

	if err != nil {
		app.Logger().Fatalf(err.Error())
		return nil
	}

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
		// ctx.ServeFile("../../web/public/index.html")
	})

	api := app.Party("/api")
	{
		mvc.New(api.Party("/signup")).Register(connectedDB).Handle(new(user.Signup))
		mvc.New(api.Party("/signin")).Register(connectedDB).Handle(new(user.Signin))
		mvc.New(api.Party("/users")).Register(connectedDB).Handle(new(user.Controller))
		mvc.New(api.Party("/apps")).Register(connectedDB).Handle(new(user.Controller))
		mvc.New(api.Party("/resources")).Register(connectedDB).Handle(new(user.Controller))
		mvc.New(api.Party("/teams")).Register(connectedDB).Handle(new(user.Controller))
	}

	// app.Get("/{p:path}", func(ctx iris.Context) {
	// 	ctx.ServeFile("../../web/public/index.html")
	// })

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fmt.Println("basepath")
	fmt.Println(b)
	fmt.Println(basepath)

	app.HandleDir("/", iris.Dir("../../web/public"), iris.DirOptions{IndexName: "index.html"})

	return app

}

func main() {
	app := newApp()
	port := ":8081"
	if s := os.Getenv("PORT"); s != "" {
		port = ":" + s
	}

	app.Listen(port)
}
