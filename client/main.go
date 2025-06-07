package client

import (
	"fmt"
	"net/http"
)

func New() *App {
	return &App{
		routes: make(map[Route]http.HandlerFunc),
	}
}

func (app *App) AppendRoute(routeKey Route, handler http.HandlerFunc) {
	app.routes[routeKey] = handler
}

func (app *App) Get(routeName string, handler http.HandlerFunc) {
	routeKey := Route{routeName: routeName, method: Get}
	app.AppendRoute(routeKey, handler)
}

func (app *App) Start() {
	routes := app.routes

	fmt.Println("Server is running on :3000 with the following routes:")
	for route := range routes {
		fmt.Printf("- %s [%s]\n", route.routeName, route.method)
	}

	for route := range routes {
		handlerKey := route
		http.HandleFunc(route.routeName, routes[handlerKey])
	}

	http.ListenAndServe(":3000", nil)
}

// app.get("name", "middleware", handler)
