package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var A = []int{-3, 1, 2, -2, 5, 6}
	fmt.Println(strconv.Itoa(Solution(A)))

}

func Solution(A []int) int {
	sort.Ints(A)
	maxP := A[len(A)-2] * A[len(A)-1] * A[len(A)]
	return maxP

}
