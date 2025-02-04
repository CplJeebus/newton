package main

import (
	"fmt"
	"strconv"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	var A = []int{4, 2, 2, 5, 1, 1, 1, 5, 8}
	fmt.Println(strconv.Itoa(Solution(A)))
}

func Solution(A []int) int {
	var smallestA float64
	var currentA float64
	var currentS float64
	var startSlice int
	var j int
	for i := 0; i < len(A); i++ {
		currentS = float64(A[i])
		for j = i + 1; j < len(A); j++ {

			currentS = currentS + float64(A[j])
			currentA = currentS / float64((j - i + 1))
			if i == 0 && j == i+1 {
				smallestA = currentA
			}
			if currentA < smallestA {
				smallestA = currentA
				startSlice = i
				fmt.Println("Change S A " + strconv.FormatFloat(smallestA, 'f', 3, 64) + "i " + strconv.Itoa(i) + "j" + strconv.Itoa(j))
			}
			fmt.Println("S A " + strconv.FormatFloat(smallestA, 'f', 3, 64) + "i " + strconv.Itoa(i) + "j" + strconv.Itoa(j))
		}
	}
	return startSlice
}
