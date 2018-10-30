package main

import (
	bootstrap "./bootstrap"
	controllers "./controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := bootstrap.CreateApp()

	//declareRoutes(app)

	// set the view engine target to ./templates folder
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))

	mvc.New(app).Handle(new(controllers.MainController))

	app.Run(iris.Addr(":3001"), iris.WithoutServerError(iris.ErrServerClosed))
}
