package handlers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
)

type notFoundHandler struct {
}

func (n *notFoundHandler) Handle(request *httpMessage.Request) (*httpMessage.Response, error) {
	status := constants.ERROR_CODE_NOT_FOUND
	message := constants.STATUS_NOT_FOUND

	response := &httpMessage.Response{
		Status: status,
		Headers: httpMessage.Header{
			constants.HEADER_CONTENT_TYPE:   []string{"text/plain"},
			constants.HEADER_CONTENT_LENGTH: []string{"0"},
			"Connection":                    []string{"close"},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
	}
	return response, nil
}

func NewNotFoundHandler() IHandler {
	return &notFoundHandler{}
}

var _ IHandler = (*notFoundHandler)(nil)
