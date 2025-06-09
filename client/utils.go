package client

import (
	"fmt"
	"net/http"
	"strings"
)

// appendRoute adds a new route to the application's route map, checking for duplicates.
func (app *App) appendRoute(method Method, routeName, pattern string, handler HandlerFunc) {
	if app.routes[routeName] == nil {
		app.routes[routeName] = make(map[Method]Route)
	}

	if _, exists := app.routes[routeName][method]; exists {
		panic(fmt.Sprintf("Duplicate route: %s [%s]", pattern, method))
	}

	app.routes[routeName][method] = Route{
		pattern: pattern,
		handler: handler,
	}
}

// Path returns the request path.
func (ctx *Context) Path() string {
	return ctx.Request.URL.Path
}

// extractStaticPrefix returns the static prefix of a route pattern (ignoring parameters).
func extractStaticPrefix(pattern string) string {
	segments := strings.Split(strings.Trim(pattern, "/"), "/")
	var static []string

	for _, seg := range segments {
		if !strings.HasPrefix(seg, ":") {
			static = append(static, seg)
		}
	}
	return "/" + strings.Join(static, "/")
}

// routeHandler is the main HTTP handler that matches incoming requests to registered routes and executes middleware and handlers.
func (app *App) routeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := Method(r.Method)

	for _, methodMap := range app.routes {
		for m, route := range methodMap {
			if m != method {
				continue
			}

			requestSegments := strings.Split(strings.Trim(path, "/"), "/")
			patternSegments := strings.Split(strings.Trim(route.pattern, "/"), "/")

			if len(requestSegments) != len(patternSegments) {
				continue
			}

			params := make(map[string]string)
			matched := true

			for i := range patternSegments {
				if strings.HasPrefix(patternSegments[i], ":") {
					paramName := strings.TrimPrefix(patternSegments[i], ":")
					params[paramName] = requestSegments[i]
				} else if patternSegments[i] != requestSegments[i] {
					matched = false
					break
				}
			}

			if matched {
				ctx := &Context{
					Writer:  w,
					Request: r,
					param:   params,
				}

				for _, mw := range app.middlewares {
					mw(ctx)
					if ctx.terminated {
						return
					}
				}

				route.handler(ctx)
				return
			}
		}
	}

	// If no match was found
	http.NotFound(w, r)
}

// Method returns the HTTP request method (e.g. GET, POST).
func (ctx *Context) Method() string {
	return ctx.Request.Method
}

// Handle registers a new route with the given method, path, middleware and handler function
func (app *App) handle(method Method, pattern string, mw Middleware, handler HandlerFunc) {
	route := extractStaticPrefix(pattern)

	app.appendRoute(method, route, pattern, wrapMiddleware(handler, mw))
}

// wrapMiddleware wraps a handler with a middleware, ensuring the middleware runs before the handler and can terminate the chain.
func wrapMiddleware(handler HandlerFunc, mw Middleware) HandlerFunc {
	if mw == nil {
		return handler
	}

	return func(ctx *Context) {
		mw(ctx)
		if !ctx.terminated {
			handler(ctx)
		}
	}
}
