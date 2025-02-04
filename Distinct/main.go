package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	var A = []int{2, 1, 1, 2, 3, 1, 6, 6, 6, 1, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(strconv.Itoa(Solution(A)))
}

func Solution(A []int) int {
	slices.Sort(A)
	curCounter := 1
	counter := 0
	// Counter := make([]int,len(A)) // Not right for now.
	for k, _ := range A {
		if k < len(A)-1 && A[k] == A[k+1] {
			curCounter++
		} else {
			counter++
			curCounter = 1
		}
	}
	return counter
}
