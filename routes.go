package main

import "github.com/kataras/iris"

func declareRoutes(app *iris.Application) {
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Ciao amico tuo</h1>")
	})

	app.Get("/pang", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
}
