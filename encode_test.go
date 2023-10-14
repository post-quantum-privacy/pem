package pem_test

import (
	"bytes"
	"fmt"
	"testing"

	stdPem "encoding/pem"

	"github.com/post-quantum-privacy/pem"
)

func TestEncode(t *testing.T) {
	data := `MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGZNv6T7dNmLSyrC6j9C5UMKQ/Sf
rayfBzM9mhNHzV/tXBqCnTL4lQ2fHarAhVbJ2fP2nXXJWgjP1L5OxHXmcZaUUWU9
P/0cv6BZAvx4sH/CN0o+gxWgV0PTZ9tMvf94XHNc157qi9grPRkahhsv7ujRAEJ2
D6CIPBoHkmZKQlxjAgMBAAE=`

	buf := bytes.NewBufferString(data)

	output := bytes.NewBuffer(nil)
	kind := "RSA PUBLIC KEY"
	n, err := pem.Encode(kind, buf, output, nil)
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
}
