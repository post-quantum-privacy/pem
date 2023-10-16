package pem

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

var (
	ErrUnexpectedBytes = errors.New("unexpected bytes")
)

// Decode accepts a reader (which provides the PEM data) and a writer for writing the
// bytes data.
func Decode(r io.Reader, w io.Writer) (kind string, headers map[string]string, err error) {
	br := bufio.NewReader(r)
	headers = make(map[string]string, 5)

	// check for new line
	peaked, err := br.Peek(1)
	if err != nil {
		return "", nil, err
	}

	startOffset := 1
	if peaked[0] == '\n' {
		br.ReadByte()

		startOffset = 0
	}

	begin, err := br.ReadSlice(' ')
	if err != nil {
		return "", nil, err
	}

	// next read should contain "-----BEGIN". It may start
	// with a new line or not.
	if !bytes.Equal(begin, []byte(start[startOffset:])) {
		return "", nil, errors.Join(errors.New("decode begin"), ErrUnexpectedBytes)
	}

	// read kind
	kindTitle, err := br.ReadSlice('-')
	if err != nil {
		return "", nil, err
	}

	kind = string(kindTitle[:len(kindTitle)-1])

	beginEnd, err := br.ReadSlice('\n')
	if err != nil {
		return "", nil, err
	}

	// next read should contain "-----\n".
	if !bytes.Equal(beginEnd, []byte(endOfLine[1:])) {
		return "", nil, errors.Join(errors.New("decode end of line"), ErrUnexpectedBytes)
	}

	// detect headers
	peaked, err = br.Peek(65)
	if err != nil {
		if err != bufio.ErrBufferFull && err != io.EOF {
			return "", nil, err
		}

		// peaked isnt 65 chars long
	}

	// detect long header and short header
	if !bytes.Contains(peaked, []byte{'\n'}) ||
		bytes.Contains(peaked, []byte{':'}) {

		for {
			// read full line
			headerLine, _, err := br.ReadLine()
			if err != nil {
				return "", nil, err
			}

			// end of headers
			if len(headerLine) == 0 {
				break
			}

			k, v, _ := bytes.Cut(headerLine, []byte{':'})
			headers[string(k)] = string(v)
		}
	}

	// decode base64 data region and auto detect footer start
	ff := newFooterFinder(br)
	b64r := base64.NewDecoder(base64.StdEncoding, ff)

	buf := make([]byte, 1024)
	for {
		n, err := b64r.Read(buf)
		if err != nil {
			if err != io.EOF {
				return kind, headers, err
			}

			// end of file
			if n == 0 {
				break
			}
		}

		_, err = w.Write(buf[:n])
		if err != nil {
			return kind, headers, err
		}
	}

	// use ff.rest and read more if needed
	fr := io.MultiReader(bytes.NewReader(ff.rest), br)
	bfr := bufio.NewReader(fr)

	line, _, err := bfr.ReadLine()
	if err != nil {
		return kind, headers, err
	}

	if !strings.HasPrefix(string(line), end[1:]+kind+endOfLine[:5]) {
		return kind, headers, errors.Join(errors.New("decode footer"), ErrUnexpectedBytes)
	}

	return kind, headers, nil
}
