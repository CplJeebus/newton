package main

import (
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type pair struct {
	A int
	B int
}

type genePair struct {
	Name  string
	Value int
}

func main() {
	var P = []int{2, 5, 0, 7}
	var Q = []int{4, 5, 6, 13}
	fmt.Println(seqSolution("CAGCCTATTTGGGTT", P, Q))
}

func seqSolution(S string, P []int, Q []int) []int {

	var p = make([]pair, len(P))
	var rV []int
	var x = 0
	var subS string
	var curX int
	var geneSeq = []genePair{{"A", 1}, {"C", 2}, {"G", 3}, {"T", 4}}
	var curS genePair

	for i := 0; i < len(P); i++ {
		p[i] = pair{P[i], Q[i]}
	}
	var pr pair
	for _, pr = range p {
		subS = S[pr.A : pr.B+1]
		for i := 0; i < len(subS); i++ {

			for _, curS = range geneSeq {
				if string(subS[i]) == curS.Name {
					curX = curS.Value
				}
			}

			if x == 0 {
				x = curX
			}
			if x >= curX {
				x = curX
			}

		}
		rV = append(rV, x)
		x = 0
	}
	return rV
}
