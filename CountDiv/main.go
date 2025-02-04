package main

import (
	"fmt"
	"math"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//A := 23456
	//B := 345890234
	A := 56
	B := 894
	K := 3
	fmt.Println(Solution(A, B, K))

}

func Solution(A int, B int, K int) int {
	prefixS := make([]int, B-A+2)
	for i := A; i < B+1; i++ {
		j := i - A
		prefixS[j+1] = prefixS[j]
		//fmt.Println("i " + strconv.Itoa(i) + " j " + strconv.Itoa(j))
		//fmt.Println("hello " + strconv.Itoa(int(math.Mod(float64(i), float64(K)))))
		if math.Mod(float64(i), float64(K)) == 0 {
			prefixS[j+1]++
		}
	}
	//fmt.Println(prefixS)
	return prefixS[B-A+1]
}
