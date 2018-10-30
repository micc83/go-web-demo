package bootstrap

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func CreateApp() *iris.Application {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	return app
}
