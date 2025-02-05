package main

import (
	"fmt"
	"strings"
)

// max returns the larger of a or b.
func maxS(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Solution computes the maximal double slice sum in array A.
func Solution(A []int) int {
	N := len(A)
	// Arrays to hold the maximum subarray sums ending at each position from the left,
	// and starting at each position from the right.
	left := make([]int, N)
	right := make([]int, N)

	// Calculate left[i]: maximum sum of a subarray ending at i,
	// for i from 1 to N-2. We initialize with 0s since we allow empty sub-slices.
	for i := 1; i < N-1; i++ {
		left[i] = maxS(0, left[i-1]+A[i])

	}
	fmt.Println(left)
	fmt.Println(strings.Repeat("*", 25))
	// Calculate right[i]: maximum sum of a subarray starting at i,
	// for i from N-2 down to 1.
	for i := N - 2; i > 0; i-- {
		right[i] = maxS(0, right[i+1]+A[i])

	}
	fmt.Println(right)
	maxDoubleSlice := 0
	// Now, for every possible middle index Y (from 1 to N-2),
	// the candidate double slice sum is left[Y-1] + right[Y+1].
	for Y := 1; Y < N-1; Y++ {
		candidate := left[Y-1] + right[Y+1]
		fmt.Println("Y: ", Y, " Sum: ", candidate, " Left: ", left[Y-1], "Right: ", right[Y+1])
		maxDoubleSlice = maxS(maxDoubleSlice, candidate)
	}

	return maxDoubleSlice
}

func main() {
	// Example test case:
	// A = [3, 2, 6, -1, 4, 5, -1, 2]
	// Expected maximal double slice sum: 17
	A := []int{3, 2, 6, -1, 4, 5, -1, 2}
	result := Solution(A)
	fmt.Println("Maximal double slice sum is:", result) // Expected output: 17
}
