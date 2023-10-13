package pem

import (
	"io"
)

type Block struct {
	// RSA PRIVATE KEY, ECC PUBLIC KEY... etc
	Type string
	// Optional headers, can be nil
	Headers map[string]string
	// Data within the PEM block
	Data io.ReadWriter
}

const (
	start     = "\n-----BEGIN "
	end       = "\n-----END "
	endOfLine = "-----\n"
)
