// Implementation of func wordCount.
// Counts the number of times words are used in a given string.
// https://tour.golang.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	
	for _, v := range strings.Fields(s) {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	
	return
}

func main() {
	wc.Test(WordCount)
}
