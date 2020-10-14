package main

import (
	"os"

	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.Default()
	// app.Logger().SetLevel("debug")

	// app.Use(recover.New())
	// app.Use(logger.New())

	api := app.Party("/api")
	{
		api.Get("/", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "hello", "status": iris.StatusOK})
		})

	}

	app.HandleDir("/", iris.Dir("../../web/public"), iris.DirOptions{IndexName: "index.html"})

	port := ":8081"
	if s := os.Getenv("PORT"); s != "" {
		port = ":" + s
	}

	app.Listen(port)
}
