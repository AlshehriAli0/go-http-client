# Go HTTP Client

[![WIP](https://img.shields.io/badge/status-WIP-yellow.svg)](https://github.com/yourusername/go-http-client)

A lightweight HTTP client for Go inspired by Express.js patterns. This project aims to provide a familiar development experience for Express.js developers transitioning to Go, while maintaining Go's performance characteristics.

## Installation

```bash
go get github.com/ali/go-http-client
```

## Overview

This client provides an Express.js-like interface for building HTTP servers in Go, built on top of the standard `net/http` package.

## Features

- Express.js style routing
- URL parameter handling (`/:param`)
- Query parameter support
- JSON response handling
- Middleware support
- Minimal API surface
- Standard library based

## Examples

### Basic Route
```go
app.Get("/", func(ctx *client.Context) {
    ctx.Send("Hello World")
})
```

### URL Parameters
```go
// Route: "/:id" - Captures any value after the root path as 'id' parameter
// e.g. "/123" -> Returns "123"
app.Get("/:id", func(ctx *client.Context) {
    param := ctx.Param("id")
    ctx.Send(param)
})
```

### Query Parameters
```go
// Route: "/search?s=query" - Retrieves 's' query parameter from URL
// e.g. "/search?s=hello" -> Returns "hello"
app.Get("/search", func(ctx *client.Context) {
    search := ctx.Query("s")
    ctx.Send(search)
})
```

### JSON Response
```go
type Test struct {
    Name string `json:"name"`
    Job  string `json:"job"`
}

app.Get("/json", func(ctx *client.Context) {
    data := Test{Name: "Ali", Job: "dev"}
    ctx.JSON(data)
})
```

### Complete Example
```go
package main

import (
    "go-http-client/client"
)

type Test struct {
    Name string `json:"name"`
    Job  string `json:"job"`
}

func main() {
    app := client.New()
    
    // URL parameter example
    app.Get("/:id", func(ctx *client.Context) {
        param := ctx.Param("id")
        ctx.Send(param)
    })

    // Query parameter example
    app.Get("/search", func(ctx *client.Context) {
        search := ctx.Query("s")
        ctx.Send(search)
    })

    // JSON response example
    app.Get("/json", func(ctx *client.Context) {
        data := Test{Name: "Ali", Job: "dev"}
        ctx.JSON(data)
    })

    // Start the server on port 3000
    app.Start(3000)
}
```

## Status

This project is currently a work in progress and is not recommended for production use.

## License

MIT 