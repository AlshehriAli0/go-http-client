package main

import (
	"go-http-client/client"
)

type Test struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func main() {
	app := client.New()

	testJson := Test{Name: "Ali", Job: "dev"}

	app.Get("/:id", func(ctx *client.Context) {
		param := ctx.Param("id")
		ctx.Send(param)
	})

	app.Get("/json", func(ctx *client.Context) {
		ctx.JSON(testJson)
	})

	app.Start()

}
