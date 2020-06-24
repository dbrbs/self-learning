// Implementation of a function to pass to pic.Show.
// Generates an image using the returned 2D slice.
// https://tour.golang.org/moretypes/18

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := range picture {
		temp := make([]uint8, dx)
		for j := range temp {
			// Formula used is the sum of consecutive integers from i to j, inclusive
			temp[j] = uint8((i+j)*(i-j-1)/2)
		}
		picture[i] = temp
	}
	
	return picture
}

func main() {
	pic.Show(Pic)
}

