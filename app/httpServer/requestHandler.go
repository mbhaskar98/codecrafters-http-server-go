package httpServer

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
	"strings"
)

type RequestHandler struct{}

func (r *RequestHandler) HandleRequest(request httpMessage.Request) (*httpMessage.Response, error) {
	status := constants.ERROR_CODE_OK
	message := constants.STATUS_OK

	var body = new(string)
	if request.Path == nil {
		status = constants.ERROR_CODE_NOT_FOUND
		message = constants.STATUS_NOT_FOUND
	} else if strings.HasPrefix(*request.Path, "/echo/") {
		*body = strings.TrimPrefix(*request.Path, "/echo/")
	} else if *request.Path != "/" {
		status = constants.ERROR_CODE_NOT_FOUND
		message = constants.STATUS_NOT_FOUND
	}
	response := &httpMessage.Response{
		Status: status,
		Headers: httpMessage.Header{
			"Content-Type": []string{"text/plain"},
			"Connection":   []string{"close"},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    200,
	}

	if body != nil && len(*body) > 0 {
		response.Body = []byte(*body)
		response.Headers["Content-Length"] = []string{fmt.Sprintf("%d", len(*body))}
	}

	return response, nil
}

func NewResponseBuilder() *RequestHandler {
	return &RequestHandler{}
}
