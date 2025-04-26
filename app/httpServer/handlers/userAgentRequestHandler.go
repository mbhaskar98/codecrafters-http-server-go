package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
)

type userAgentRequestHandler struct {
}

func (u userAgentRequestHandler) Handle(request *httpMessage.Request) (*httpMessage.Response, error) {
	userAgent := request.Headers["User-Agent"][0]

	status := constants.ERROR_CODE_OK
	message := constants.STATUS_OK
	body := userAgent

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

func NewUserAgentRequestHandler() IHandler {
	return &userAgentRequestHandler{}
}

var _ IHandler = (*userAgentRequestHandler)(nil)
