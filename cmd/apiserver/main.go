package main

import (
	"os"

	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.Default()
	app.Logger().SetLevel("debug")

	// app.Use(recover.New())
	// app.Use(logger.New())

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	app.Logger().Fatalf("connect to sqlite3 failed")
	// 	return
	// }

	// iris.RegisterOnInterrupt(func() {
	// 	defer db.Close()
	// })

	// if os.Getenv("ENV") != "" {
	// 	db.DropTableIfExists(&User{}) // drop table
	// }
	// db.AutoMigrate(&User{}) // create table: // AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
		// ctx.ServeFile("../../web/public/index.html")
	})

	api := app.Party("/api")
	{
		api.Get("/register", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
		})

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

	app.HandleDir("/", iris.Dir("../../web/public"), iris.DirOptions{IndexName: "index.html"})

	port := ":8081"
	if s := os.Getenv("PORT"); s != "" {
		port = ":" + s
	}

	app.Listen(port)
}
