package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}

	for i, s := range p {
		switch {
		case s >= 'A' && s <= 'M':
			p[i] = s + 13
		case s >= 'N' && s <= 'Z':
			p[i] = s - 13
		case s >= 'a' && s <= 'm':
			p[i] = s + 13
		case s >= 'n' && s <= 'z':
			p[i] = s - 13
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
