package handlers

import (
	"fmt"
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
			"Content-Type": []string{"text/plain"},
			"Connection":   []string{"close"},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
	}
	response.Headers["Content-Length"] = []string{fmt.Sprintf("0")}
	return response, nil
}

func NewNotFoundHandler() IHandler {
	return &notFoundHandler{}
}

var _ IHandler = (*notFoundHandler)(nil)
