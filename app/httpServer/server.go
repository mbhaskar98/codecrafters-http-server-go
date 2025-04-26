package httpServer

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/parser"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/router"
	"net"
)

type Server struct {
	listener net.Listener
	parser   parser.IParser
	router   router.IRouter
}

func (server *Server) Run() {
	for {
		data, conn, err := server.getRequest()
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
		fmt.Println(httpResponse.ToString())
		_, err = server.sendResponse([]byte(httpResponse.ToString()), conn)
		if err != nil {
			fmt.Println("Error sending response: ", err.Error())
			return
		}

	}
}

func (server *Server) getRequest() ([]byte, net.Conn, error) {
	conn, err := server.listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		return nil, nil, err
	}
	bytes := make([]byte, 1024)
	_, err = conn.Read(bytes)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		return nil, nil, err
	}
	return bytes, conn, nil
}

func (server *Server) sendResponse(response []byte, conn net.Conn) (int, error) {
	n, err := conn.Write(response)

	if err != nil {
		fmt.Println("Error writing to connection: ", err.Error())
		return -1, err
	}
	err = conn.Close()
	if err != nil {
		fmt.Println("Error closing connection: ", err.Error())
	}
	return n, err
}

func NewServer(parser parser.IParser, router router.IRouter) (*Server, error) {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		return nil, err
	}
	return &Server{
		listener: l,
		parser:   parser,
		router:   router,
	}, nil
}
