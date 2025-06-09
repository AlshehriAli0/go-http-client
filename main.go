package main

import (
	"fmt"

	"github.com/AlshehriAli0/go-http-client/client"
)

// User represents a simple user data structure
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Logger is a middleware that prints request information to the console.
func Logger(ctx *client.Context) {
	fmt.Printf("New request: %s %s\n", ctx.Request.Method, ctx.Request.URL.Path)
}

// main is the entry point of the application.
func main() {
	// Create a new app instance
	app := client.New()

	// Add global middleware
	app.Use(Logger)

	// Example 1: Simple GET route
	app.Get("/", nil, func(ctx *client.Context) {
		ctx.Send("Welcome to the API!")
	})

	// Example 2: Route with URL parameter
	app.Get("/users/:id", nil, func(ctx *client.Context) {
		userID := ctx.Param("id")
		ctx.Send(fmt.Sprintf("Fetching user with ID: %s", userID))
	})

	// Example 3: Search query params with JSON response
	app.Get("/profile", nil, func(ctx *client.Context) {
		query := ctx.Query("name")

		user := User{
			Name:  query,
			Email: fmt.Sprintf("%s@example.com", query),
		}
		ctx.JSON(user)
	})

	// Example 4: POST request
	app.Post("/users", nil, func(ctx *client.Context) {
		body, err := ctx.ReadBody()
		if err != nil {
			ctx.Error("Failed to read request body", 400)
			return
		}
		ctx.Send(fmt.Sprintf("Received data: %s", body))
	})

	// Start the server on port 3000
	app.Start(3000)
}
