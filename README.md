# Go HTTP Client

[![WIP](https://img.shields.io/badge/status-WIP-yellow.svg)](https://github.com/yourusername/go-http-client)

A lightweight HTTP client for Go inspired by Express.js patterns. This project aims to provide a familiar development experience for Express.js developers transitioning to Go, while maintaining Go's performance characteristics.

## Overview

This client provides an Express.js-like interface for building HTTP servers in Go, built on top of the standard `net/http` package.

## Features

- Express.js-style routing
- Middleware support
- Minimal API surface
- Standard library based

## Example

```go
package main

import (
    "fmt"
    "go-http-client/client"
    "net/http"
)

func main() {
    app := client.New()

    app.Get("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello")
    })

    app.Start()
}
```

## Installation

```bash
go get github.com/yourusername/go-http-client
```

## Status

This project is currently a work in progress and is not recommended for production use.

## License

MIT 