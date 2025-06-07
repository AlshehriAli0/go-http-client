package client

import "net/http"

func New() *App {
	return &App{
		routes: make(map[Route]http.HandlerFunc),
	}
}

func (app *App) AppendRoute(routeKey Route, handler http.HandlerFunc) {
	app.routes[routeKey] = handler
}

func (app *App) Get(routeName string, middleware func(), handler http.HandlerFunc) {
	routeKey := Route{routeName: routeName, method: "get"}
	app.AppendRoute(routeKey, handler)
}

// app.get("name", "middleware", handler)
