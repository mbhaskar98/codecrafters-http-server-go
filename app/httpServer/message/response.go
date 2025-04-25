package httpMessage

import (
	"fmt"
	"strings"
)

type Response struct {
	Status  int
	Body    []byte
	Headers Header
	Version string
	Message string
	Reason  string
	Code    int
}

func (r *Response) ToString() string {
	var response strings.Builder
	response.WriteString(fmt.Sprintf("%s %d %s\r\n", r.Version, r.Status, r.Message))

	for key, values := range r.Headers {
		for _, value := range values {
			response.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
		}
	}

	response.WriteString("\r\n")
	response.Write(r.Body)

	return response.String()
}
