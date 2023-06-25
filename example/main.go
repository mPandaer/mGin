package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("welcome my gee!\n"))
	})

	engine.GET("/info", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "path: %s\n host: %s\n", request.URL.Path, request.Host)
	})

	engine.Run(":8080")
}
