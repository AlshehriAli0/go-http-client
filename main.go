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
