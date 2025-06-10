package main

import (
	"fmt"

	client "github.com/AlshehriAli0/go-http-client"
)

// AuthMiddleware is a simple example middleware for demonstration.
func AuthMiddleware(ctx *client.Context) {
	token := ctx.GetHeader("X-Token")
	if token != "secret" {
		ctx.Status(401)
		ctx.Send("Unauthorized")
		ctx.End()
	}
}

func UserMiddleware(ctx *client.Context) {
	println("This only runs on all routes inside /user group after the global middleware and before route specific middleware")
}

func main() {
	app := client.New()

	// Global middleware
	app.Use(func(ctx *client.Context) {
		fmt.Printf("[LOG] %s %s\n", ctx.Method(), ctx.Path())
	})

	// Group for user-related routes
	userGroup := app.Group("/users", UserMiddleware)

	// Now all routes will start with /users
	userGroup.Get("/", nil, func(ctx *client.Context) {
		ctx.Send("List all users")
	})
	userGroup.Get("/:id", nil, func(ctx *client.Context) {
		ctx.Send("Get user with ID: " + ctx.Param("id"))
	})
	userGroup.Post("/", AuthMiddleware, func(ctx *client.Context) {
		ctx.Send("Create user (auth required)")
	})

	// Group for admin-related routes
	adminGroup := app.Group("/admin")
	adminGroup.Get("/dashboard", AuthMiddleware, func(ctx *client.Context) {
		ctx.Send("Welcome to the admin dashboard!")
	})

	app.Start(3000)
}
