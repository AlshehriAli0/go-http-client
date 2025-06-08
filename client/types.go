package client

import "net/http"

type Method string

const (
	Post   Method = "post"
	Get    Method = "get"
	Patch  Method = "patch"
	Update Method = "update"
	Delete Method = "delete"
)

type RouteEntry struct {
	pattern   string
	routeName string
	method    Method
}

type App struct {
	routes map[RouteEntry]HandlerFunc
}

type HandlerFunc func(*Context)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	param   map[string]string
}
