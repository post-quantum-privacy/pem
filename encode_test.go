package pem_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/post-quantum-privacy/pem"
)

func TestEncode(t *testing.T) {
	buf := bytes.NewBufferString(`MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGZNv6T7dNmLSyrC6j9C5UMKQ/Sf
rayfBzM9mhNHzV/tXBqCnTL4lQ2fHarAhVbJ2fP2nXXJWgjP1L5OxHXmcZaUUWU9
P/0cv6BZAvx4sH/CN0o+gxWgV0PTZ9tMvf94XHNc157qi9grPRkahhsv7ujRAEJ2
D6CIPBoHkmZKQlxjAgMBAAE=`)

	output := bytes.NewBuffer(nil)
	n, err := pem.Encode("RSA PUBLIC KEY", buf, output, nil)
	if err != nil {
		t.Error(err)
	}

	if n != output.Len() {
		t.Errorf("n does not equal output len. n = %d, real = %d", n, output.Len())
	}

	fmt.Println(output.String())
}
