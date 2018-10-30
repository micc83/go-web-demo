package main

import (
	bootstrap "./bootstrap"
	"github.com/kataras/iris"
)

func main() {
	app := bootstrap.CreateApp()

	declareRoutes(app)

	app.Run(iris.Addr(":3001"), iris.WithoutServerError(iris.ErrServerClosed))
}
