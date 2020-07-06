// Implementation of square root approximation.
// https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
)

func Sqrt(x float64) (z float64) {
	z = 1.0
	
	// prints guess before performing next update
	for i := 0; i < 10; i += 1 {
		fmt.Println(z)
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}

