package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
)

type defaultRequestHandler struct {
}

func (d defaultRequestHandler) Handle(request *httpMessage.Request) (*httpMessage.Response, error) {
	status := constants.ERROR_CODE_OK
	message := constants.STATUS_OK

	response := &httpMessage.Response{
		Status: status,
		Headers: httpMessage.Header{
			"Content-Type": []string{"text/plain"},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
	}
	response.Headers["Content-Length"] = []string{fmt.Sprintf("0")}
	return response, nil
}

func NewDefaultRequestHomeHandler() IHandler {
	return &defaultRequestHandler{}
}

var _ IHandler = (*defaultRequestHandler)(nil)
