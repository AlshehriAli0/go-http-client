package client

import "strings"

// appendRoute adds a new route to the application's route map, checking for duplicates
func (app *App) appendRoute(routeKey RouteEntry, handler HandlerFunc) {
	// Check route duplication
	for existingRoute := range app.routes {
		if existingRoute.routeName == routeKey.routeName {
			panic("There are two routes with the same signature")
		}
	}
	app.routes[routeKey] = handler
}

// generateRouteKey creates a new RouteEntry from a route string and HTTP method
func generateRouteKey(route string, method Method) RouteEntry {
	parsedRouteName := strings.Split(route, ":")[0]
	routeKey := RouteEntry{routeName: parsedRouteName, method: method, pattern: route}

	return routeKey

}

// Handle registers a new route with the given method, path, middleware and handler function
func (app *App) handle(method Method, route string, mw Middleware, handler HandlerFunc) {
	key := generateRouteKey(route, method)
	app.appendRoute(key, wrapMiddleware(handler, mw))
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
