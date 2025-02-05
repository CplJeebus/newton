package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	inBytes, _ := os.Open("./names.txt")
	var namesList []string
	var nameElements []string

	elements := flag.Bool("e", false, "Enable the names list")
	namesC := flag.Bool("n", false, "Enable the names list")

	flag.Parse()

	scan := bufio.NewScanner(inBytes)
	for scan.Scan() {
		S := standardise(scan.Text())
		namesList = append(namesList, S)
	}
	// count := 0
	sort.Strings(namesList)
	// fmt.Println(namesList)
	switch {
	case *namesC:
		namMap := countName(namesList)
		for k, v := range namMap {
			fmt.Printf("Name: %s Count: %d\n", k, v)
		}
	case *elements:
		namMap := countName(namesList)
		for k, _ := range namMap {
			for _, word := range splitName(k) {
				nameElements = append(nameElements, word)
			}
		}
		sort.Strings(nameElements)
		countElements(nameElements)

	}
}

func splitName(name string) []string {
	return strings.Split(name, " ")
}

func countElements(names []string) map[string]int {
	elements := make(map[string]int)
	for _, name := range names {
		elements[name]++
	}
	for k, v := range elements {
		fmt.Printf("Element: %s Count: %d\n", k, v)
	}
	return elements
}

func countName(names []string) map[string]int {
	namMap := make(map[string]int)
	for _, name := range names {
		namMap[name]++
	}
	return namMap
}

func standardise(S string) string {
	S = strings.Replace(S, "\"", "", -1)
	S = strings.Trim(S, " ")
	S = strings.ToLower(S)
	return S
}
