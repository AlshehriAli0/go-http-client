package client

import "net/http"

type RouteEntry struct {
	pattern   string
	routeName string
	method    Method
}

type Middleware func(*Context)

type App struct {
	routes      map[RouteEntry]HandlerFunc
	middlewares []Middleware
}

type HandlerFunc func(*Context)

type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	param      map[string]string
	terminated bool
}

type Method string

const (
	Post   Method = "post"
	Get    Method = "get"
	Patch  Method = "patch"
	Update Method = "update"
	Delete Method = "delete"
)
