package httpServer

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
	"strings"
)

type RequestParser struct {
}

func (p *RequestParser) Parse(request []byte) (*httpMessage.Request, error) {
	// Split headers and body
	parts := bytes.SplitN(request, []byte("\r\n\r\n"), 2)
	if len(parts) < 1 {
		return nil, errors.New("invalid request: no header section")
	}

	headerSection := string(parts[0])
	bodySection := []byte{}
	if len(parts) == 2 {
		bodySection = parts[1]
	}

	lines := strings.Split(headerSection, "\r\n")
	if len(lines) < 1 {
		return nil, errors.New("invalid request: empty request line")
	}

	// Parse request line
	requestLine := strings.Fields(lines[0])
	if len(requestLine) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", lines[0])
	}

	method := requestLine[0]
	path := requestLine[1]
	version := requestLine[2]

	// Parse headers
	headers := make(httpMessage.Header)
	for _, line := range lines[1:] {
		colonIdx := strings.Index(line, ":")
		if colonIdx == -1 {
			continue // skip malformed header lines
		}
		key := strings.TrimSpace(line[:colonIdx])
		value := strings.TrimSpace(line[colonIdx+1:])
		headers[key] = append(headers[key], value)
	}

	// Return parsed Request
	return &httpMessage.Request{
		Version: &version,
		Method:  &method,
		Path:    &path,
		Headers: headers,
		Body:    &bodySection,
	}, nil
}

func NewRequestParser() *RequestParser {
	return &RequestParser{}
}
