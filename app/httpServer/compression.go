package httpServer

import (
	"fmt"
)

type SupportedCompressionTechniques string

const (
	GZIP SupportedCompressionTechniques = "gzip"
)

func Compress(data []byte, technique SupportedCompressionTechniques) ([]byte, error) {
	switch technique {
	case GZIP:
		return data, nil
	default:
		return data, fmt.Errorf("unsupported compression technique: %s", technique)
	}
}
