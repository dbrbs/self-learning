// Implements rot13Reader.
// https://tour.golang.org/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// helper function for rotating
func rotate(b byte) (byte) {
	if ('A' <= b && b <= 'M') || ('a' <= b && b <= 'm') {
		return b + 13
	} else if ('N' <= b && b <= 'Z') || ('n' <= b && b <= 'z') {
		return b - 13
	} else {
		return b
	}
}

// Read function for reader
func (r13 rot13Reader) Read(b13 []byte) (int, error) {
	n := len(b13)
	b := make([]byte, n)
	
	i,e := r13.r.Read(b)
		
	for j,v := range b {
		b13[j] = rotate(v)
	}
	
	return i,e
}		

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

