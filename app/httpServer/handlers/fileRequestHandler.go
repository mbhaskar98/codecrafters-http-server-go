package handlers

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/constants"
	"github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type fileRequestHandler struct {
	fileDirectory string
}

func (f *fileRequestHandler) Handle(request *httpMessage.Request) (*httpMessage.Response, error) {
	fileToServe := strings.TrimPrefix(*request.Path, "/files/")
	if len(fileToServe) == 0 {
		return nil, fmt.Errorf("invalid file path: %s", fileToServe)
	}

	fullPath := filepath.Join(f.fileDirectory, fileToServe)

	// Read file
	content, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	status := constants.ERROR_CODE_OK
	message := constants.STATUS_OK

	response := &httpMessage.Response{
		Status: status,
		Headers: httpMessage.Header{
			"Content-Type":   []string{"application/octet-stream"},
			"Content-Length": []string{strconv.Itoa(len(content))},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
		Body:    content,
	}

	return response, nil
}

func NewFileRequestHandler(fileDirectory string) IHandler {
	return &fileRequestHandler{
		fileDirectory: fileDirectory,
	}
}

var _ IHandler = (*fileRequestHandler)(nil)
