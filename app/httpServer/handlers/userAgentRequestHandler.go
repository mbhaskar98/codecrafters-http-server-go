package handlers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
	"strconv"
)

type userAgentRequestHandler struct {
}

func (u *userAgentRequestHandler) Handle(request *httpMessage.Request) (*httpMessage.Response, error) {
	userAgent := request.Headers[constants.HEADER_USER_AGENT][0]

	status := constants.ERROR_CODE_OK
	message := constants.STATUS_OK
	body := userAgent

	response := &httpMessage.Response{
		Status: status,
		Headers: httpMessage.Header{
			constants.HEADER_CONTENT_TYPE:   []string{"text/plain"},
			constants.HEADER_CONTENT_LENGTH: []string{strconv.Itoa(len(body))},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
		Body:    []byte(body),
	}

	return response, nil
}

func NewUserAgentRequestHandler() IHandler {
	return &userAgentRequestHandler{}
}

var _ IHandler = (*userAgentRequestHandler)(nil)
