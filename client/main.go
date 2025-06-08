package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

func New() *App {
	return &App{
		routes: make(map[RouteEntry]HandlerFunc),
	}
}

// Routes
func (app *App) Get(route string, handler HandlerFunc) {
	key := generateRouteKey(route, Get)
	app.appendRoute(key, handler)
}

func (app *App) Post(route string, handler HandlerFunc) {
	key := generateRouteKey(route, Post)
	app.appendRoute(key, handler)
}

func (app *App) Update(route string, handler HandlerFunc) {
	key := generateRouteKey(route, Update)
	app.appendRoute(key, handler)
}

func (app *App) Patch(route string, handler HandlerFunc) {
	key := generateRouteKey(route, Patch)
	app.appendRoute(key, handler)
}

func (app *App) Delete(route string, handler HandlerFunc) {
	key := generateRouteKey(route, Delete)
	app.appendRoute(key, handler)
}

// Context methods
func (ctx *Context) Send(str string) {
	strByte := []byte(str)
	ctx.Writer.Write(strByte)

}

func (ctx *Context) Param(paramKey string) string {
	return ctx.Params[paramKey]
}

func (ctx *Context) Query(queryParam string) string {
	query := ctx.Request.URL.Query()
	return query.Get(queryParam)
}

func (ctx *Context) JSON(jsonData interface{}) {
	parsedJson, err := json.Marshal(jsonData)

	if err != nil {
		http.Error(ctx.Writer, "Not a valid JSON", http.StatusInternalServerError)
		return
	}

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.Write(parsedJson)
}

func (app *App) Start() {
	routes := app.routes

	fmt.Println("Server is running on :3000 with the following routes:")
	for route := range routes {
		fmt.Printf("- %s [%s]\n", route.pattern, route.method)
	}

	for route := range routes {
		routeKey := route
		handler := routes[routeKey]

		http.HandleFunc(routeKey.routeName, func(w http.ResponseWriter, r *http.Request) {
			params := make(map[string]string)

			paramKey := strings.Split(routeKey.pattern, "/:")[1]
			paramValue := path.Base(r.URL.Path)

			params[paramKey] = paramValue

			ctx := &Context{
				Writer:  w,
				Request: r,
				Params:  params,
			}

			handler(ctx)
		})
	}

	http.ListenAndServe(":3000", nil)
}

// app.get("name", "middleware", handler)
