package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13Transform(b byte) byte {
	b_rot := b + 13
	if (b >= 65) && (b <= 90) {
		if b_rot > 90 {
			return (b_rot - 90) + 64
		}
		return b_rot
	} else if (b >= 97) && (b <= 122) {
		if b_rot > 122 {
			return (b_rot - 122) + 96
		}
		return b_rot
	}
	return b
}

func (rotReader rot13Reader) Read(b []byte) (int, error) {
	raw := make([]byte, len(b))
	n, err := rotReader.r.Read(raw)
	if err == nil {
		for i, char := range raw {
			// fmt.Println(char, rot13Transform(char))
			b[i] = rot13Transform(char)
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
