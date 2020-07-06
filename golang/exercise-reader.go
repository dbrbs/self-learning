// Implements a Reader for MyReader.
// https://tour.golang.org/methods/22

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Infinite stream of 'A'
func (r MyReader) Read(b []byte) (int, error) {
	for i:=0; i < 8; i++ {
		b[i] = 'A'
	}
	return 8, nil
}

func main() {
	reader.Validate(MyReader{})
}

