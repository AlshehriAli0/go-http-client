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

// Group creates a new route group with the specified prefix. Useful for modular route organization.
func (app *App) Group(prefix string, mws ...Middleware) *Group {
	return &Group{prefix: normalizePath(prefix), app: app, middlewares: mws}
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
