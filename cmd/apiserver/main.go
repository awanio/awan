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

	// app.Use(recover.New())
	// app.Use(logger.New())

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	sqlDB, err := db.DB()

	if err != nil {
		app.Logger().Fatalf("connect to sqlite3 failed")
		return nil
	}

	iris.RegisterOnInterrupt(func() {
		defer sqlDB.Close()
	})

	// if os.Getenv("ENV") != "" {
	// 	db.DropTableIfExists(&user.Users{}) // drop table
	// }
	db.AutoMigrate(&user.Users{})
	db.AutoMigrate(&user.Credentials{})

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
		// ctx.ServeFile("../../web/public/index.html")
	})

	api := app.Party("/api")
	{
		mvc.New(app.Party("/signup")).Handle(new(user.Controller))

		// api.Get("/signup", func(ctx iris.Context) {
		// 	ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
		// })

		api.Get("/login", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
		})

		api.Get("/users", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
		})

		api.Get("/apps", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
		})

		api.Get("/resources", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
		})

		api.Get("/teams", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
		})

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
