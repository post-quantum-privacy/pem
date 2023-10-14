package pem

import (
	"bytes"
	"io"
)

type footerFinder struct {
	r    io.Reader
	rest []byte
}

func newFooterFinder(r io.Reader) *footerFinder {
	return &footerFinder{
		r: r,
	}
}

func (f *footerFinder) Read(p []byte) (n int, err error) {
	n, err = f.r.Read(p)
	if err != nil {
		if err != io.EOF {
			return n, err
		}
	}

	// does it contain the ending of the PEM
	if s := bytes.Index(p, []byte{'-'}); s != -1 {
		f.rest = p[s:]
		p = p[:s]

		return len(p), io.EOF
	}

	return n, err
}
