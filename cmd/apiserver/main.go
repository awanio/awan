package main

import (
	"os"

	"github.com/awanio/awan/internal/env"
	"github.com/awanio/awan/internal/runtime"
	"github.com/awanio/awan/internal/user"
	"github.com/awanio/awan/pkg/helper"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func newApp() *iris.Application {

	app := iris.Default()
	app.Logger().SetLevel("debug")

	runtime.Setup()

	if runtime.DBerror != nil {
		app.Logger().Fatalf("connect to sqlite3 failed")
		return nil
	}

	iris.RegisterOnInterrupt(func() {
		defer runtime.SQLDB.Close()
	})

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
		// ctx.ServeFile("../../web/public/index.html")
	})

	err := runtime.DB.AutoMigrate(&user.Users{}, &user.Credentials{})

	if err != nil {
		app.Logger().Fatalf(err.Error())
		return nil
	}

	userRepository := user.NewRepository(runtime.DB)
	// Generate admin user
	admin, creationStatus, err := userRepository.CreateAdmin()

	if err != nil {
		app.Logger().Fatalf(err.Error())
		return nil
	}

	if creationStatus {
		app.Logger().Info("admin password: ", admin["adminPassword"])
		app.Logger().Info("admin username: ", admin["adminUsername"])
	}

	// cors middleware
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: false,
	})

	app.UseGlobal(crs)

	// exclude from auth middleware
	mvc.New(app.Party("/api/signin")).Register(userRepository).Handle(new(user.Signin))

	var mySecret = []byte(env.JWTSecret)

	j := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	api := app.Party("/api")
	{
		api.Use(j.Serve)

		mvc.New(api.Party("/users")).Register(userRepository).Handle(new(user.Controller))
		mvc.New(api.Party("/apps")).Register(runtime.DB).Handle(new(user.Controller))
		mvc.New(api.Party("/resources")).Register(runtime.DB).Handle(new(user.Controller))
		mvc.New(api.Party("/teams")).Register(runtime.DB).Handle(new(user.Controller))
	}

	// app.Get("/{p:path}", func(ctx iris.Context) {
	// 	ctx.ServeFile("../../web/public/index.html")
	// })

	app.HandleDir("/", iris.Dir(helper.FromBasepath("web/public")), iris.DirOptions{IndexName: "index.html"})

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
