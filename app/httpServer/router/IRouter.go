package router

import "github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"

type IRouter interface {
	HandleRequest(request *httpMessage.Request) (*httpMessage.Response, error)
}
