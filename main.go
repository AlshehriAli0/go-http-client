package main

import (
	"fmt"
	"go-http-client/client"
	"net/http"
)

// Test struct represents a simple data structure with Name and Job fields
type Test struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

// Logger is a middleware that prints request method and path
func Logger(ctx *client.Context) {
	// here you can check auth or log etc..
	fmt.Println("Request:", ctx.Request.Method, ctx.Request.URL.Path)
}

// Auth is a middleware that checks authentication and redirects to /json if not authenticated
func Auth(ctx *client.Context) {
	// here you can check auth or log etc..
	fmt.Println("User Not Authenticated, Redirecting")
	ctx.Redirect("/json", http.StatusPermanentRedirect)
}

// main initializes the application and sets up example routes
func main() {
	app := client.New()

	// Register middleware
	app.Use(Logger)

	// Create test JSON data
	testJson := Test{Name: "Ali", Job: "dev"}

	// Example 1: URL Parameter handling
	// Route: "/:id" - Captures any value after the root path as 'id' parameter
	// e.g. "/123" -> Returns "123"
	app.Get("/:id", nil, func(ctx *client.Context) {
		param := ctx.Param("id")
		ctx.Send(param)
	})

	// Example 2: Query Parameter handling
	// Route: "/search?s=query" - Retrieves 's' query parameter from URL
	// e.g. "/search?s=hello" -> Returns "hello"
	app.Get("/search", nil, func(ctx *client.Context) {
		search := ctx.Query("s")
		ctx.Send(search)
	})

	// Example 3: JSON Response
	// Route: "/json" - Returns a JSON object
	// Returns: {"name":"Ali","job":"dev"}
	app.Get("/json", nil, func(ctx *client.Context) {
		ctx.JSON(testJson)
	})

	// Example 4: POST request handling with authentication
	// Route: "/post" - Accepts POST requests with body data
	// Requires authentication via Auth middleware
	app.Post("/post", Auth, func(ctx *client.Context) {
		body, err := ctx.ReadBody()
		if err != nil {
			ctx.Error("Error reading body", http.StatusBadRequest)
		}
		ctx.Send(body)
	})

	// Start the server on port 3000
	app.Start(3000)
}
