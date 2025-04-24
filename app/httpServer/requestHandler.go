package httpServer

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
)

type RequestHandler struct{}

func (r *RequestHandler) HandleRequest(request HttpRequest) (*HttpResponse, error) {
	status := constants.ERROR_CODE_OK
	message := constants.STATUS_OK

	if request.Path != nil && *request.Path != "/" {
		status = constants.ERROR_CODE_NOT_FOUND
		message = constants.STATUS_NOT_FOUND
	}
	var body *string = nil
	response := &HttpResponse{
		Status: status,
		//Body:   []byte("Hello World"),
		//Headers: Header{
		//	"Content-Type": []string{"text/plain"},
		//  "Connection": []string{"close"},
		//  "Connection": []string{"close"}
		//},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    200,
	}

	if body != nil && len(*body) > 0 {
		response.Headers["Content-Length"] = []string{fmt.Sprintf("%d", len(*body))}
	}

	return response, nil
}

func NewResponseBuilder() *RequestHandler {
	return &RequestHandler{}
}
