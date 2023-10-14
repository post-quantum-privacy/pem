package pem_test

import (
	"bytes"
	"fmt"
	"testing"

	stdPem "encoding/pem"

	"github.com/post-quantum-privacy/pem"
)

func TestEncode(t *testing.T) {
	data := []byte(`MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGZNv6T7dNmLSyrC6j9C5UMKQ/Sf
rayfBzM9mhNHzV/tXBqCnTL4lQ2fHarAhVbJ2fP2nXXJWgjP1L5OxHXmcZaUUWU9
P/0cv6BZAvx4sH/CN0o+gxWgV0PTZ9tMvf94XHNc157qi9grPRkahhsv7ujRAEJ2
D6CIPBoHkmZKQlxjAgMBAAE=`)

	validateEncode(t, "RSA PUBLIC KEY", data, nil)
}

func TestEncodeHeaders(t *testing.T) {
	data := []byte(`MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGZNv6T7dNmLSyrC6j9C5UMKQ/Sf
rayfBzM9mhNHzV/tXBqCnTL4lQ2fHarAhVbJ2fP2nXXJWgjP1L5OxHXmcZaUUWU9
P/0cv6BZAvx4sH/CN0o+gxWgV0PTZ9tMvf94XHNc157qi9grPRkahhsv7ujRAEJ2
D6CIPBoHkmZKQlxjAgMBAAE=`)

	headers := make(map[string]string)
	headers["hash"] = "sha256"

	validateEncode(t, "RSA PUBLIC KEY", data, headers)

	headers["Proc-Type"] = "something"
	headers["kdf"] = "sha256"
	headers["cipher_suite"] = "aes-256-gcm"
	validateEncode(t, "RSA PUBLIC KEY", data, headers)
}

func validateEncode(t *testing.T, kind string, data []byte, headers map[string]string) {
	output := bytes.NewBuffer(nil)

	n, err := pem.Encode(kind, bytes.NewBuffer(data), output, headers)
	if err != nil {
		t.Error(err)
	}

	if n != output.Len() {
		t.Errorf("n does not equal output len. n = %d, real = %d", n, output.Len())
	}

	fmt.Println(output.String())

	p, rest := stdPem.Decode(output.Bytes())
	if len(rest) != 0 {
		t.Errorf("rest is %d, should be zero", len(rest))
	}

	// they do equal, likely a issue with carrige returns
	if !bytes.Equal([]byte(data), p.Bytes) {
		t.Error("data bytes do not match")
	}

	if p.Type != kind {
		t.Error("kind/type does not match")
	}

	if len(p.Headers) != len(headers) {
		t.Error("not all headers have been written")
	}

	for k, v := range headers {
		if p.Headers[k] != v {
			t.Errorf("header %q does not have the correct value", k)
		}
	}
}
