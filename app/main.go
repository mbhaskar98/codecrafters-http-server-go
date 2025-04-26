package main

import (
	"flag"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/handlers"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/parser"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/router"
	"os"
)

var (
	directory = flag.String("directory", ".", "File directory")
)

func main() {
	flag.Parse()
	config := httpServer.Config{
		Directory: *directory,
	}

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	p := parser.NewRequestParser()
	r := router.NewRouter(
		httpServer.GenerateRoutes(config),
		handlers.NewDefaultRequestHomeHandler(),
		handlers.NewNotFoundHandler(),
	)
	server, err := httpServer.NewServer(p, r, config)
	if err != nil {
		fmt.Println("Failed to create server")
		os.Exit(1)
	}
	fmt.Println("Server created")
	server.Run()
}
