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
	switch *request.Method {
	case constants.HTTP_METHOD_GET:
		return f.handleGetRequest(request)
	case constants.HTTP_METHOD_POST:
		return f.handlePostRequest(request)
	}
	return nil, fmt.Errorf("unsupported method: %s", *request.Method)
}

func (f *fileRequestHandler) handleGetRequest(request *httpMessage.Request) (*httpMessage.Response, error) {
	fileToServe := f.getFilePath(request)
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
			constants.HEADER_CONTENT_TYPE:   []string{"application/octet-stream"},
			constants.HEADER_CONTENT_LENGTH: []string{strconv.Itoa(len(content))},
		},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
		Body:    content,
	}

	return response, nil
}

func (f *fileRequestHandler) handlePostRequest(request *httpMessage.Request) (*httpMessage.Response, error) {
	fileToServe := f.getFilePath(request)
	if len(fileToServe) == 0 {
		return nil, fmt.Errorf("invalid file path: %s", fileToServe)
	}

	contentLength := request.Headers[constants.HEADER_CONTENT_LENGTH][0]
	contentLengthInt, err := strconv.Atoi(contentLength)
	if err != nil {
		return nil, fmt.Errorf("failed to parse content length: %w", err)
	}

	if contentLengthInt > 0 {
		fullPath := filepath.Join(f.fileDirectory, fileToServe)
		err = os.WriteFile(fullPath, (*request.Body)[:contentLengthInt], 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to write file: %w", err)
		}
	}
	status := constants.ERROR_CODE_CREATED
	message := constants.STATUS_CREATED

	response := &httpMessage.Response{
		Status:  status,
		Headers: httpMessage.Header{},
		Version: constants.HTTP_VERSION_1_1,
		Message: message,
		Reason:  message,
		Code:    status,
	}

	return response, nil

}

func (f *fileRequestHandler) getFilePath(request *httpMessage.Request) string {
	return strings.TrimPrefix(*request.Path, "/files/")
}

func NewFileRequestHandler(fileDirectory string) IHandler {
	return &fileRequestHandler{
		fileDirectory: fileDirectory,
	}
}

var _ IHandler = (*fileRequestHandler)(nil)
