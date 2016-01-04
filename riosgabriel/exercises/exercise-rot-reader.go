package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)

	for i := 0; i < n; i++ {
		aux := b[i]

		switch {
			case 'a' <= aux && aux <= 'z':
				if aux > 'm' {
					b[i] -= 13
				} else {
					b[i] += 13
				}

			case (aux >= 'A' && aux <= 'Z'):
				if aux > 'M' {
					b[i] -= 13
				} else {
					b[i] += 13
				}
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}