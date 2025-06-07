package client

import "net/http"

func (app *App) appendRoute(routeKey Route, handler http.HandlerFunc) {
	app.routes[routeKey] = handler
}
