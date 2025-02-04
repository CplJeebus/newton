package main

import (
	"fmt"
	"strconv"
)

// Solution returns an array of minimal impact factors for each query.
func Solution(S string, P []int, Q []int) []int {
	N := len(S)
	M := len(P)

	// Create prefix sum arrays for A, C and G.
	// We don't need T because if none of A, C, G occur, then T must be present.
	prefixA := make([]int, N+1)
	prefixC := make([]int, N+1)
	prefixG := make([]int, N+1)

	// Build prefix sums: for i from 0 to N-1, fill prefix arrays at i+1.
	for i := 0; i < N; i++ {
		// Carry over previous counts.
		prefixA[i+1] = prefixA[i]
		prefixC[i+1] = prefixC[i]
		prefixG[i+1] = prefixG[i]

		switch S[i] {
		case 'A':
			prefixA[i+1]++
		case 'C':
			prefixC[i+1]++
		case 'G':
			prefixG[i+1]++
			// For T we don't need an array; if none of the above increase then T is present.
		}
	}
	tmp := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(tmp)
	fmt.Println(prefixA)
	fmt.Println(prefixC)
	fmt.Println(prefixG)

	// Allocate result slice.
	result := make([]int, M)

	// Process each query.
	for i := 0; i < M; i++ {
		// Define the query range as [from, to] (inclusive). We use to+1 for prefix array indexing.
		from := P[i]
		to := Q[i] + 1

		// Check in order: A (impact factor 1), C (2), G (3). If none of these are present, then answer is T (4).
		if prefixA[to]-prefixA[from] > 0 {
			fmt.Println("A " + strconv.Itoa(prefixA[to]-prefixA[from]))
			result[i] = 1
		} else if prefixC[to]-prefixC[from] > 0 {
			fmt.Println("C " + strconv.Itoa(prefixC[to]-prefixC[from]))
			result[i] = 2
		} else if prefixG[to]-prefixG[from] > 0 {
			fmt.Println("G " + strconv.Itoa(prefixG[to]-prefixG[from]))
			result[i] = 3
		} else {
			fmt.Println("T ")
			result[i] = 4
		}
	}

	return result
}

func main() {
	// Test example: S = "CAGCCTA"
	// Queries:
	// P[0] = 2, Q[0] = 4 → substring "GCC" → minimal impact factor is 2 (for C)
	// P[1] = 5, Q[1] = 5 → substring "T"   → impact factor is 4
	// P[2] = 0, Q[2] = 6 → whole string  → minimal impact factor is 1 (for A)
	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}
	result := Solution(S, P, Q)
	fmt.Println(result) // Expected output: [2, 4, 1]
}
