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

// Logger middleware prints request information
func Logger(ctx *client.Context) {
	fmt.Printf("New request: %s %s\n", ctx.Request.Method, ctx.Request.URL.Path)
}

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

	// Example 3: Route with query parameters
	app.Get("/search", nil, func(ctx *client.Context) {
		query := ctx.Query("q")
		ctx.Send(fmt.Sprintf("Searching for: %s", query))
	})

	// Example 4: JSON response
	app.Get("/profile", nil, func(ctx *client.Context) {
		user := User{
			Name:  "John Doe",
			Email: "john@example.com",
		}
		ctx.JSON(user)
	})

	// Example 5: POST request
	app.Post("/users", nil, func(ctx *client.Context) {
		body, err := ctx.ReadBody()
		if err != nil {
			ctx.Error("Failed to read request body", 400)
			return
		}
		ctx.Send(fmt.Sprintf("Received data: %s", body))
	})

	// Start the server on port 3000
	fmt.Println("Server starting on http://localhost:3000")
	app.Start(3000)
}
