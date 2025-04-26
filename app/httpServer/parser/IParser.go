package parser

import "github.com/codecrafters-io/http-server-starter-go/app/httpServer/httpMessage"

type IParser interface {
	Parse(request []byte) (*httpMessage.Request, error)
}
