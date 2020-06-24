// Implements fibonacci function using a closure.
// https://tour.golang.org/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a,b := 1,0
	return func() int {
		b = b + a
		a = b - a
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

