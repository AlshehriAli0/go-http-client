package client

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// New creates and returns a new App instance with initialized routes.
func New() *App {
	return &App{
		routes: make(map[RouteEntry]HandlerFunc),
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
	for route := range routes {
		fmt.Printf("- %s [%s]\n", route.pattern, route.method)
	}

	for route := range routes {
		routeKey := route
		handler := routes[routeKey]

		http.HandleFunc(routeKey.routeName, func(w http.ResponseWriter, r *http.Request) {

			param := make(map[string]string)

			// Check if route has params
			if strings.Contains(routeKey.pattern, "/:") {
				paramKey := strings.Split(routeKey.pattern, "/:")[1]
				paramValue := path.Base(r.URL.Path)
				param[paramKey] = paramValue
			}

			ctx := &Context{
				Writer:  w,
				Request: r,
				param:   param,
			}

			for _, mw := range app.middlewares {
				mw(ctx)
				// stop the chaining
				if ctx.terminated {
					return
				}
			}

			handler(ctx)
		})
	}

	portStr := strconv.Itoa(port)
	if err := http.ListenAndServe(":"+portStr, nil); err != nil {
		log.Fatal(err)
	}
}

// app.get("name", "middleware", handler)
