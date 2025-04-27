package handlers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
)

type defaultRequestHandler struct {
}

func (d *defaultRequestHandler) Handle(request *httpMessage.Request) (*httpMessage.Response, error) {
	status := constants.ERROR_CODE_OK
	message := constants.STATUS_OK

	response := &httpMessage.Response{
		Status: status,
		Headers: httpMessage.Header{
			constants.HEADER_CONTENT_TYPE:   []string{"text/plain"},
			constants.HEADER_CONTENT_LENGTH: []string{"0"},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
	}
	return response, nil
}

func NewDefaultRequestHomeHandler() IHandler {
	return &defaultRequestHandler{}
}

var _ IHandler = (*defaultRequestHandler)(nil)
