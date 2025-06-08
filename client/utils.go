package client

import "strings"

func (app *App) appendRoute(routeKey RouteEntry, handler HandlerFunc) {
	// Check route duplication
	for existingRoute := range app.routes {
		if existingRoute.routeName == routeKey.routeName {
			panic("There are two routes with the same signature")
		}
	}
	app.routes[routeKey] = handler
}

func generateRouteKey(route string, method Method) RouteEntry {
	parsedRouteName := strings.Split(route, ":")[0]
	routeKey := RouteEntry{routeName: parsedRouteName, method: method, pattern: route}

	return routeKey

}

func (app *App) Handle(method Method, route string, mw Middleware, handler HandlerFunc) {
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
