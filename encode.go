package pem

import (
	"bufio"
	"encoding/base64"
	"errors"
	"io"
	"sort"
	"strings"
)

var (
	ErrHeaderColon = errors.New("pem header key cannot include a colon")
)

func Encode(kind string, data io.Reader, w io.Writer, headers map[string]string) (n int, err error) {
	o := bufio.NewWriter(w)

	if err := validateHeaders(headers); err != nil {
		return n, err
	}

	// write "BEGINS..."
	wn, err := o.WriteString(start[1:])
	n += wn
	if err != nil {
		return n, err
	}

	wn, err = o.WriteString(kind + endOfLine)
	n += wn
	if err != nil {
		return n, err
	}

	// write headers
	if headers != nil || len(headers) > 0 {
		headersList := make([]string, 0, len(headers))

		for k, v := range headers {
			// Proc-Type must be written first
			if k == "Proc-Type" {
				wn, err := o.WriteString("Proc-Type" + ":" + v)
				n += wn
				if err != nil {
					return n, err
				}

				continue
			}

			headersList = append(headersList, k)
		}

		// headers should be a consistent order
		sort.Strings(headersList)
		for i := 0; i <= len(headersList); i++ {
			// write header
			wn, err := o.WriteString(headersList[i] + ":" + headers[headersList[i]])
			n += wn
			if err != nil {
				return n, err
			}
		}

		wn, err := o.WriteRune('\n')
		n += wn
		if err != nil {
			return n, err
		}
	}

	// encode data and break lines
	b64w := base64.NewEncoder(base64.StdEncoding, o)
	buf := make([]byte, 48)

	for {
		rn, err := data.Read(buf)
		if err != nil {
			if err != io.EOF {
				return n, err
			} else if rn == 0 {
				break
			}
		}

		wn, err = b64w.Write(buf[:rn])
		n += int(float64(wn) * 1.351)
		if err != nil {
			return n, err
		}

		// break line on 64th char
		if wn == 48 {
			wn, err = o.WriteRune('\n')
			n += wn
			if err != nil {
				return n, err
			}
		}
	}

	if err := b64w.Close(); err != nil {
		return n, err
	}

	// write "ENDS..."
	wn, err = o.WriteString(end + kind + endOfLine)
	n += wn
	if err != nil {
		return n, err
	}

	return n, o.Flush()
}

// validateHeaders checks if there is a ":" within the keys
func validateHeaders(headers map[string]string) error {
	for k := range headers {
		if strings.Contains(k, ":") {
			return ErrHeaderColon
		}
	}

	return nil
}
