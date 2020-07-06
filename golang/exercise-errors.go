// Implements error-handling for Sqrt function.
// https://tour.golang.org/methods/20

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// Error function: returns error string
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

// Square root function
func Sqrt(x float64) (float64, error) {	
	// check if negative number
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	
	// prints guess before performing next update
	for i := 0; i < 10; i += 1 {
		fmt.Println(z)
		z -= (z*z - x) / (2*z)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

