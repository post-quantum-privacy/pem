# PEM

[![Go Report Card](https://goreportcard.com/badge/post-quantum-privacy/pem)](https://goreportcard.com/report/post-quantum-privacy/pem)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/post-quantum-privacy/pem)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/post-quantum-privacy/pem/.github/workflows/go.yml)

PEM is a standards complient encoder for the PEM format. This libary enables the streaming of PEM encodes, something which is not supported within Go's standard libary. This enables the encoding and decoding of large PEM files without exhusting system memory.

## The Problem This Resolves
Go's standard `encoding/pem` package does not allow the streaming of files to be encoded or decoded, thus making large files exhust the device's RAM.
**This libarary allows data to be encoded and decoded in a stream** (io.Reader & io.Writer), without commiting large buffers to memory, or worse yet,
the entire file. 

## Import

```
go get github.com/post-quantum-privacy/pem
```

## Example

```go
package main

import (
	"bytes"
    "fmt"

    "github.com/post-quantum-privacy/pem"
)

func main() {
    // the data you whish to encode with PEM
    data := bytes.NewBuffer([]byte("..."))
    // the output writter. Could be a file, conn or any io.Writer
    out := bytes.NewBuffer(nil)

    // encode the message
    n, err := pem.Encode("PGP MESSAGE", data, out, nil)
	if err != nil {
		panic(err)
	}

    fmt.Println(out.String())
}
```

```go
package main

import (
	"bytes"
    "fmt"

    "github.com/post-quantum-privacy/pem"
)

func main() {
    // the data you whish to encode with PEM
    data := bytes.NewBuffer([]byte("..."))
    // the output writter. Could be a file, conn or any io.Writer
    out := bytes.NewBuffer(nil)

    // define your headers
    headers := make(map[string]string)
    headers["hash"] = "SHA256"

    // encode the message
    n, err := pem.Encode("PGP MESSAGE", data, out, headers)
	if err != nil {
		panic(err)
	}

    fmt.Println(out.String())
}
```