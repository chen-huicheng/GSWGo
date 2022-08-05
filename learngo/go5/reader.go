package main

import (
	"io"
	"unicode"
)

type MyReader struct{}

func (m MyReader) Read(out []byte) (int, error) {
	for i, _ := range out {
		out[i] = 'A'
	}
	return len(out), nil
}

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(out []byte) (n int, err error) {
	n, err = r13.r.Read(out)
	for i := range out {
		if unicode.IsLetter(rune(out[i])) {
			continue
		}
		if out[i] >= 'a' && out[i] <= 'z' {
			out[i] = out[i] - 'a' + 13
			out[i] = out[i]%26 + 'a'

		} else {
			out[i] = out[i] - 'A' + 13
			out[i] = out[i]%26 + 'A'
		}
	}
	return
}

func main() {
	reader.Validate(Myreader{})
}
