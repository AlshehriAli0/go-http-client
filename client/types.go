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

type Route struct {
	routeName string
	method    Method
}

type App struct {
	routes map[Route]HandlerFunc
}

type HandlerFunc func(*Context)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}
