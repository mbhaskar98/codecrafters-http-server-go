package httpServer

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/parser"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/router"
	"net"
)

type Config struct {
	Directory string
}

type Server struct {
	listener net.Listener
	parser   parser.IParser
	router   router.IRouter
	config   Config
}

func (server *Server) Run() {
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go server.serveRequest(conn)
	}
}

func (server *Server) serveRequest(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection: ", err.Error())
		}
	}(conn)

	data, err := server.getData(conn)
	if err != nil {
		fmt.Println("Error getting request: ", err.Error())
		return
	}

	httpRequest, err := server.parser.Parse(data)
	if err != nil {
		fmt.Println("Error parsing request: ", err.Error())
		return
	}
	httpResponse, err := server.router.HandleRequest(httpRequest)
	if err != nil {
		fmt.Println("Error handling request: ", err.Error())
		return
	}
	_, err = server.sendResponse([]byte(httpResponse.ToString()), conn)
	if err != nil {
		fmt.Println("Error sending response: ", err.Error())
		return
	}
}

func (server *Server) getData(conn net.Conn) ([]byte, error) {
	bytes := make([]byte, 1024)
	_, err := conn.Read(bytes)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		return nil, err
	}
	return bytes, nil
}

func (server *Server) sendResponse(response []byte, conn net.Conn) (int, error) {
	n, err := conn.Write(response)

	if err != nil {
		fmt.Println("Error writing to connection: ", err.Error())
		return -1, err
	}
	return n, err
}

func NewServer(
	parser parser.IParser,
	router router.IRouter,
	config Config,
) (*Server, error) {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		return nil, err
	}
	return &Server{
		listener: l,
		parser:   parser,
		router:   router,
		config:   config,
	}, nil
}
