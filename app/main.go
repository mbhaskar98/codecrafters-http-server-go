package main

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	parser := httpServer.NewRequestParser()
	requestHandler := httpServer.NewResponseBuilder()
	server, err := httpServer.NewServer(parser, requestHandler)
	if err != nil {
		fmt.Println("Failed to create server")
		os.Exit(1)
	}
	fmt.Println("Server created")
	server.Run()
}
