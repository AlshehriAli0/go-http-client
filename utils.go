package client

import (
	"fmt"
	"net/http"
	"strings"
)

func normalizePath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if len(path) > 1 && strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	return path
}

// Path returns the request path.
func (ctx *Context) Path() string {
	return ctx.Request.URL.Path
}

func (app *App) routeHandler(w http.ResponseWriter, r *http.Request) {
	path := normalizePath(r.URL.Path)
	method := Method(r.Method)

	for pattern, methods := range app.routes {
		route, ok := methods[method]
		if !ok {
			continue
		}

		patternSegments := strings.Split(strings.Trim(pattern, "/"), "/")
		pathSegments := strings.Split(strings.Trim(path, "/"), "/")

		if len(patternSegments) != len(pathSegments) {
			continue
		}

		params := make(map[string]string)
		matched := true

		for i := range patternSegments {
			if strings.HasPrefix(patternSegments[i], ":") {
				key := strings.TrimPrefix(patternSegments[i], ":")
				params[key] = pathSegments[i]
			} else if patternSegments[i] != pathSegments[i] {
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

	http.NotFound(w, r)
}

// Method returns the HTTP request method (e.g. GET, POST).
func (ctx *Context) Method() string {
	return ctx.Request.Method
}

func (app *App) handle(method Method, pattern string, mw Middleware, handler HandlerFunc) {
	pattern = normalizePath(pattern)

	if app.routes[pattern] == nil {
		app.routes[pattern] = make(map[Method]Route)
	}

	if _, exists := app.routes[pattern][method]; exists {
		panic(fmt.Sprintf("Duplicate route: %s [%s]", pattern, method))
	}

	app.routes[pattern][method] = Route{
		pattern: pattern,
		handler: wrapMiddleware(handler, mw),
	}
}

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
