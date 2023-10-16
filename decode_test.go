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

	if len(headers) > 0 {
		t.Error("headers should be empty")
	}

	if kind != "RSA PRIVATE KEY" {
		t.Error("kind is incorrect")
	}

	if !bytes.Equal(data, body.Bytes()) {
		t.Error("bytes do not match")
	}
}

func TestDecodeHeaders(t *testing.T) {
	data := []byte("some secret key here")
	out := bytes.NewBuffer(nil)
	headers := make(map[string]string)
	headers["hash"] = "SHA256"
	headers["type"] = "ecdsa"
	headers["key"] = "X448"

	_, err := pem.Encode("PGP MESSAGE", bytes.NewBuffer(data), out, headers)
	if err != nil {
		t.Fatal(err)
	}

	body := bytes.NewBuffer(nil)
	kind, headers, err := pem.Decode(out, body)
	if err != nil {
		t.Fatal(err)
	}

	if len(headers) != 3 {
		t.Error("headers should equal 3")
	}

	if kind != "PGP MESSAGE" {
		t.Error("kind is incorrect")
	}

	if !bytes.Equal(data, body.Bytes()) {
		t.Error("bytes do not match")
	}
}
