package handlers

import "github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"

type IHandler interface {
	Handle(request *httpMessage.Request) (*httpMessage.Response, error)
}
