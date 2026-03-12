package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"url_shortner/internal/http"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Address (Default: localhost): ")
	address, _ := reader.ReadString('\n')
	address = strings.TrimSpace(address)
	if address == "" {
		address = "localhost"
	}

	fmt.Print("Port (Default: 8080): ")
	portStr, _ := reader.ReadString('\n')
	portStr = strings.TrimSpace(portStr)

	port, err := strconv.Atoi(portStr)
	if err != nil || port <= 0 {
		port = 8080
	}

	server, err := http.NewServer(address, port)
	if err != nil {
		return
	}
	server.Start()
}
