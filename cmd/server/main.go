package main

import (
	"url_shortner/internal/http"
)

func main() {
	server, err := http.NewServer("127.0.0.1", 8080)
	if err != nil {
		return
	}
	server.Start()
}
