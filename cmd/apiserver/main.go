package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/awanio/awan/internal/user"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newApp() *iris.Application {

	app := iris.Default()
	app.Logger().SetLevel("debug")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	sqlDB, err := db.DB()

	if err != nil {
		app.Logger().Fatalf("connect to sqlite3 failed")
		return nil
	}

	iris.RegisterOnInterrupt(func() {
		defer sqlDB.Close()
	})

	err = db.AutoMigrate(&user.Users{}, &user.Credentials{})

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
		mvc.New(api.Party("/signup")).Register(db).Handle(new(user.Signup))
		mvc.New(api.Party("/signin")).Handle(new(user.Signin))
		mvc.New(api.Party("/users")).Handle(new(user.Controller))
		mvc.New(api.Party("/apps")).Handle(new(user.Controller))
		mvc.New(api.Party("/resources")).Handle(new(user.Controller))
		mvc.New(api.Party("/teams")).Handle(new(user.Controller))
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
