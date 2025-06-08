package client

func (app *App) appendRoute(routeKey Route, handler HandlerFunc) {
	app.routes[routeKey] = handler
}
