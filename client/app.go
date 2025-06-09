package client

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// New creates and returns a new App instance with initialized routes.
func New() *App {
	return &App{
		routes: make(map[string]map[Method]Route),
	}
}

// Use adds a middleware to the application's middleware chain.
func (app *App) Use(mw Middleware) {
	app.middlewares = append(app.middlewares, mw)
}

// Routes

// Get registers a new GET route with the given path and handler function.
func (app *App) Get(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Get, route, mw, handler)
}

// Post registers a new POST route with the given path and handler function.
func (app *App) Post(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Post, route, mw, handler)
}

// Update registers a new UPDATE route with the given path and handler function.
func (app *App) Update(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Update, route, mw, handler)
}

// Patch registers a new PATCH route with the given path and handler function.
func (app *App) Patch(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Patch, route, mw, handler)
}

// Delete registers a new DELETE route with the given path and handler function.
func (app *App) Delete(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Delete, route, mw, handler)
}

// Start begins the HTTP server on the specified port and sets up all registered routes.
// It prints all registered routes to the console for debugging purposes.
func (app *App) Start(port int) {
	routes := app.routes
	if port == 0 {
		port = 3000
	}

	fmt.Printf("Server is running on :%d with the following routes:\n", port)
	for _, routeNames := range routes {
		for method, route := range routeNames {
			fmt.Printf("- %s [%s]\n", route.pattern, method)
		}
	}

	// Single entry point handles all routes
	http.HandleFunc("/", app.routeHandler)

	portStr := strconv.Itoa(port)
	if err := http.ListenAndServe(":"+portStr, nil); err != nil {
		log.Fatal(err)
	}
}
