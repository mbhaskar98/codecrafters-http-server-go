package router

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/handlers"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
	"strings"
)

type router struct {
	routes          map[string]handlers.IHandler
	notFoundHandler handlers.IHandler
	defaultHandler  handlers.IHandler
}

func (r *router) HandleRequest(request *httpMessage.Request) (*httpMessage.Response, error) {
	// TODO: Use better routing algorithm, e.g. trie
	for prefix, handler := range r.routes {
		if strings.HasPrefix(*request.Path, prefix) {
			response, err := handler.Handle(request)
			if err != nil {
				fmt.Println("Error handling request: ", err.Error())
				return r.notFoundHandler.Handle(request)
			}
			return response, nil
		}
	}
	if *request.Path == "/" {
		return r.defaultHandler.Handle(request)
	}
	return r.notFoundHandler.Handle(request)
}

func NewRouter(
	routes map[string]handlers.IHandler,
	defaultHandler handlers.IHandler,
	notFoundHandler handlers.IHandler,
) IRouter {
	return &router{
		routes:          routes,
		notFoundHandler: notFoundHandler,
		defaultHandler:  defaultHandler,
	}
}
