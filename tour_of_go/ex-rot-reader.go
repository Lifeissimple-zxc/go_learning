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
	// Read for the interface
	n, err = rot.r.Read(b)
	// This checks for EOF, think of it as a default when working wiuth such implementations
	if err != nil {
		return n, err
	}
	// iterate over chars, if it's an alphabetic character, move it by 13 nums
	for i := 0; i < n; i++ {
		ch := b[i]
		// Using unicode.IsLetter does not work bc we don't wrap around the alphabet
		switch {
		case 'A' <= ch && ch <= 'Z':
			// See how many letters we are from A & wrap around the alphabet
			b[i] = (ch-'A'+13)%26 + 'A'
		case 'a' <= ch && ch <= 'z':
			b[i] = (ch-'a'+13)%26 + 'a'
		}
	}
	return n, nil

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
