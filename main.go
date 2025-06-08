package main

import (
	"go-http-client/client"
)

// Test struct represents a simple data structure with Name and Job fields
type Test struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func main() {
	app := client.New()

	// Create test JSON data
	testJson := Test{Name: "Ali", Job: "dev"}

	// Example 1: URL Parameter handling
	// Route: "/:id" - Captures any value after the root path as 'id' parameter
	// e.g. "/123" -> Returns "123"
	app.Get("/:id", func(ctx *client.Context) {
		param := ctx.Param("id")
		ctx.Send(param)
	})

	// Example 2: Query Parameter handling
	// Route: "/search?s=query" - Retrieves 's' query parameter from URL
	// e.g. "/search?s=hello" -> Returns "hello"
	app.Get("/search", func(ctx *client.Context) {
		search := ctx.Query("s")
		ctx.Send(search)
	})

	// Example 3: JSON Response
	// Route: "/json" - Returns a JSON object
	// Returns: {"name":"Ali","job":"dev"}
	app.Get("/json", func(ctx *client.Context) {
		ctx.JSON(testJson)
	})

	// Start the server on port 3000
	app.Start(3000)

}
