package handlers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
	"strconv"
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
			"Content-Type":   []string{"text/plain"},
			"Content-Length": []string{strconv.Itoa(len(body))},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
		Body:    []byte(body),
	}

	return response, nil
}

func NewEchoRequestHandler() IHandler {
	return &echoRequestHandler{}
}

var _ IHandler = (*echoRequestHandler)(nil)
