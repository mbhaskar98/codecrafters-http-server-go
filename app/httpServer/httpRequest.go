package httpServer

import (
	"fmt"
	"strings"
)

type HttpRequest struct {
	Version *string
	Method  *string
	Path    *string
	Body    *[]byte
	Headers Header
}

func (msg *HttpRequest) ToString() string {
	var builder strings.Builder

	// Write request line
	if msg.Method != nil && msg.Path != nil && msg.Version != nil {
		builder.WriteString(fmt.Sprintf("%s %s %s\r\n", *msg.Method, *msg.Path, *msg.Version))
	}

	// Write headers
	for key, values := range msg.Headers {
		for _, value := range values {
			builder.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
		}
	}

	builder.WriteString("\r\n") // empty line between headers and body

	// Write body
	if msg.Body != nil {
		builder.Write(*(msg.Body))
	}

	return builder.String()
}
