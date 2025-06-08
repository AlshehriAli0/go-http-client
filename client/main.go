package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func New() *App {
	return &App{
		routes: make(map[Route]HandlerFunc),
	}
}

// Routes
func (app *App) Get(routeName string, handler HandlerFunc) {
	routeKey := Route{routeName: routeName, method: Get}
	app.appendRoute(routeKey, handler)
}

// context methods
func (ctx *Context) Send(str string) {
	strByte := []byte(str)
	ctx.Writer.Write(strByte)
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
		fmt.Printf("- %s [%s]\n", route.routeName, route.method)
	}

	for route := range routes {
		routeKey := route
		handler := routes[routeKey] // your custom handler

		http.HandleFunc(routeKey.routeName, func(w http.ResponseWriter, r *http.Request) {
			ctx := &Context{
				Writer:  w,
				Request: r,
			}
			handler(ctx)
		})
	}

	http.ListenAndServe(":3000", nil)
}

// app.get("name", "middleware", handler)
