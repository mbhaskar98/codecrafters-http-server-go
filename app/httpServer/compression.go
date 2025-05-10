package httpServer

import (
	"fmt"
)

type SupportedEncodingMethods string

const (
	GZIP SupportedEncodingMethods = "gzip"
)

var supportedEncodingMethods = map[SupportedEncodingMethods]bool{
	GZIP: true,
}

func Compress(data []byte, method string) ([]byte, error) {
	m := SupportedEncodingMethods(method)
	if ok := supportedEncodingMethods[m]; !ok {
		return data, fmt.Errorf("unsupported compression technique: %s", m)
	}
	switch m {
	case GZIP:
		return data, nil
	default:
		return data, fmt.Errorf("unsupported compression technique: %s", m)
	}
}
