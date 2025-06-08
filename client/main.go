package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// New creates and returns a new App instance with initialized routes
func New() *App {
	return &App{
		routes: make(map[RouteEntry]HandlerFunc),
	}
}

// Use adds a middleware to the application's middleware chain
func (app *App) Use(mw Middleware) {
	app.middlewares = append(app.middlewares, mw)
}

// Routes

// Get registers a new GET route with the given path and handler function
func (app *App) Get(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Get, route, mw, handler)
}

// Post registers a new POST route with the given path and handler function
func (app *App) Post(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Post, route, mw, handler)

}

// Update registers a new UPDATE route with the given path and handler function
func (app *App) Update(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Update, route, mw, handler)
}

// Patch registers a new PATCH route with the given path and handler function
func (app *App) Patch(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Patch, route, mw, handler)

}

// Delete registers a new DELETE route with the given path and handler function
func (app *App) Delete(route string, mw Middleware, handler HandlerFunc) {
	app.handle(Delete, route, mw, handler)

}

// Context methods

// End marks the context as terminated, stopping further middleware execution
func (ctx *Context) End() {
	ctx.terminated = true
}

// ReadBody reads and returns the request body as a string
func (ctx *Context) ReadBody() (string, error) {
	defer ctx.Request.Body.Close()

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Redirect performs an HTTP redirect to the specified route with the given status code
func (ctx *Context) Redirect(route string, code int) {
	http.Redirect(ctx.Writer, ctx.Request, route, code)
	ctx.End()
}

// Send writes a string response to the client
func (ctx *Context) Send(str string) {
	strByte := []byte(str)
	ctx.Writer.Write(strByte)

}

// Error sends an HTTP error response with the given message and status code
func (ctx *Context) Error(message string, code int) {
	http.Error(ctx.Writer, message, code)
}

// Param retrieves the value of a URL parameter by its key
// Example: For route "/users/:id", Param("id") returns the actual id value
func (ctx *Context) Param(paramKey string) string {
	return ctx.param[paramKey]
}

// Query retrieves the value of a query parameter from the URL
// Example: For URL "/search?q=test", Query("q") returns "test"
func (ctx *Context) Query(queryParam string) string {
	query := ctx.Request.URL.Query()
	return query.Get(queryParam)
}

// JSON marshals the provided data into JSON and sends it as response
// Sets appropriate Content-Type header automatically
func (ctx *Context) JSON(jsonData interface{}) {
	parsedJson, err := json.Marshal(jsonData)

	if err != nil {
		ctx.Error("Not a valid JSON", http.StatusInternalServerError)
		return
	}

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Write(parsedJson)
}

// Start begins the HTTP server on port 3000 and sets up all registered routes
// Prints all registered routes to console for debugging purposes
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
