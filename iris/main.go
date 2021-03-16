package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Get("/welcome", func(ctx iris.Context) {
		firstname := ctx.URLParamDefault("firstname", "Guest")
		lastname := ctx.URLParam("lastname")

		ctx.Writef("Hello %s %s", firstname, lastname)
	})

	app.Run(iris.Addr(":8080"))
}
