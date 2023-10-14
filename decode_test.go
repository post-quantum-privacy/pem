package pem_test

import (
	"bytes"
	"testing"

	"github.com/post-quantum-privacy/pem"
)

func TestDecode(t *testing.T) {
	data := []byte("some secret key here")
	out := bytes.NewBuffer(nil)

	_, err := pem.Encode("RSA PRIVATE KEY", bytes.NewBuffer(data), out, nil)
	if err != nil {
		t.Fatal(err)
	}

	body := bytes.NewBuffer(nil)
	kind, headers, err := pem.Decode(out, body)
	if err != nil {
		t.Fatal(err)
	}

	if headers != nil {
		t.Error("headers should be nil")
	}

	if kind != "RSA PRIVATE KEY" {
		t.Error("kind is incorrect")
	}

	if !bytes.Equal(data, body.Bytes()) {
		t.Error("bytes do not match")
	}
}
