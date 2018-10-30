package controllers

import "github.com/kataras/iris/mvc"

// MainController ...
type MainController struct{}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *MainController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

// GetView serves
// Method:   GET
// Resource: http://localhost:8080/view
func (c *MainController) GetView() mvc.Result {
	return mvc.View{
		Name: "index",
	}
}

// GetPing serves
// Method:   GET
// Resource: http://localhost:8080/ping
func (c *MainController) GetPing() string {
	return "pong"
}

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *MainController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}
