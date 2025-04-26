package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
	"strings"
)

type echoRequestHandler struct {
}

func (e *echoRequestHandler) Handle(request *httpMessage.Request) (*httpMessage.Response, error) {
	status := constants.ERROR_CODE_OK
	message := constants.STATUS_OK
	body := strings.TrimPrefix(*request.Path, "/echo/")

	response := &httpMessage.Response{
		Status: status,
		Headers: httpMessage.Header{
			"Content-Type": []string{"text/plain"},
			"Connection":   []string{"close"},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
	}
	if len(body) > 0 {
		response.Body = []byte(body)
		response.Headers["Content-Length"] = []string{fmt.Sprintf("%d", len(body))}
	}

	return response, nil
}

func NewEchoRequestHandler() IHandler {
	return &echoRequestHandler{}
}

var _ IHandler = (*echoRequestHandler)(nil)
