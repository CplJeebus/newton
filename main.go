package main

import (
	"fmt"

	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	A int
	B int
}

type genePair struct {
	Name  string
	Value int
}

func main() {
	http.HandleFunc("/newton.png", getNewton)
	http.HandleFunc("/", simple)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}

func getNewton(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := ioutil.ReadFile("newton.png")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func simple(w http.ResponseWriter, r *http.Request) {
	u := os.Getenv("NEWTON_URL")
	if u == "" {
		u = "http://localhost:8080/api"
	}
	// bare := strings.ReplaceAll(u, "api", "")

	head := `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <title>Exercises</title>
  <style>
    /* Just a bit of basic styling for clarity */
    section {
      border: 1px solid #ccc;
      padding: 1rem;
      margin: 1rem 0;
      border-radius: 4px;
    }

    h2 {
      margin-top: 0;
    }

    label {
      display: inline-block;
      width: 500px;
      margin-right: 0.5rem;
    }

    input[type="text"] {
      margin-bottom: 0.5rem;
    }

    .submit-btn {
      margin-top: 0.5rem;
      display: block;
    }
  </style>
</head> <body>`

	top := `
	<h1> <img src="newton.png" width="125"> Project Newton </h1>  
	`
	gap := `
  <section>
<h2> Given a string of 1s and 0s find the longest gap of 0s'</h2> 
  <form action="` + u + `" method="POST">
    <label for="binaryInput">Enter a string of 0s and 1s:</label>
<br> 
    <input
      type="text"
      id="gap"
      name="binaryString"
      pattern="[0-1]+"
      
      placeholder="e.g. 001101"
    />
    <button type="submit">Submit</button>
</form>
</section> `
	frogTwo := `
<section>
<h2> Frog exercise</h2> 
<form action="` + u + `" method="POST">

    <label for="binaryInput">Enter a comma seperated array of ints and an key</label>
	<br> 
    <input
      type="text"
      id="1"
      name="frogTwo"
      placeholder="e.g. a,b,c,d,e"
    />
	<input
      type="number"
      id="2"
      name="frogKey"
    />
    <button type="submit">Submit</button>
</section> `

	rotate := `
<section>
<h2> Rotate an array by a key number</h2> 
<form action="` + u + `" method="POST">

    <label for="binaryInput">Enter a comma seperated array of stings and an index</label>
	<br> 
    <input
      type="text"
      id="1"
      name="rotate"
      placeholder="e.g. a,b,c,d,e"
    />
	<input
      type="number"
      id="2"
      name="key"
      
    />
    <button type="submit">Submit</button>
</section> `

	odd := `
  <section>
<h2> Find the odd element in an array</h2> 
  <form action="` + u + `" method="POST">

    <label for="Numbers">Enter a list of numbers:</label>
<br> 
    <input
      type="text"
      id="1"
      name="Numbers"
      placeholder="0,1,2,3"
    />
    <button type="submit">Submit</button>
</form>
</section>`

	missing := `
<section>
<h2> Find the missing element in an array</h2> 
  <form action="` + u + `" method="POST">

    <label for="Numbers">Enter a list of ranging for 1 to N+1 omit one number</label>
<br> 
    <input
      type="text"
      id="1"
      name="missing"
      placeholder="2,1,3,5"
    />
    <button type="submit">Submit</button>
</form>
</section> `

	tape := ` 
<section>
<h2> Find the Equilibrium </h2> 
  <form action="` + u + `" method="POST">

    <label for="Numbers">Enter a list of numbers/label>
<br> 
    <input
      type="text"
      id="1"
      name="tape"
      placeholder="2,1,3,5"
    />
    <button type="submit">Submit</button>
</form>
</section> `

	frog := `<section>
<h2> Calculate the number of steps for a frog to cover Y </h2> 
  <form action="` + u + `" method="POST">

    <label for="Frog">Enter values for X, Y, D</label>
<br> 
    <input
      type="number"
      id="1"
      name="frog-x"
      placeholder="Vale for X"
    />
    <input
      type="number"
      id="1"
      name="frog-d"
      placeholder="Vale for D"
    />
    <input
      type="number"
      id="1"
      name="frog-y"
      placeholder="Value for Y"
    />
    <button type="submit">Submit</button>
</form>
</section>`

	tail := `
</body>
</html>`

	var b string
	b = head +
		top +
		frogTwo +
		gap +
		rotate +
		odd +
		frog +
		missing +
		tape +
		tail
	fmt.Fprint(w, b)
}

func api(w http.ResponseWriter, r *http.Request) {
	var b string
	_ = r.ParseForm()
	if r.Form.Get("binaryString") != "" {
		b = gap(r.Form.Get("binaryString"))
	} else if r.Form.Get("rotate") != "" {
		var j, _ = strconv.Atoi(r.Form.Get("key"))
		b = rotate(r.Form.Get("rotate"), j)
	} else if r.Form.Get("Numbers") != "" {
		b = oddOne(r.Form.Get("Numbers"))
	} else if r.Form.Get("frog-x") != "" {
		var x, _ = strconv.Atoi(r.Form.Get("frog-x"))
		var d, _ = strconv.Atoi(r.Form.Get("frog-d"))
		var y, _ = strconv.Atoi(r.Form.Get("frog-y"))
		b = frog(x, d, y)
	} else if r.Form.Get("missing") != "" {

		b = missing(r.Form.Get("missing"))
	} else if r.Form.Get("tape") != "" {

		b = tapeEquilibrium(r.Form.Get("tape"))
	} else if r.Form.Get("frogTwo") != "" {
		var j, _ = strconv.Atoi(r.Form.Get("frogKey"))
		b = frogTwo(r.Form.Get("frogTwo"), j)
	}
	fmt.Fprint(w, b)
}

func frog(x int, d int, y int) string {
	var result string
	var c int = 0
	for p := x; p < y; p = p + d {
		c++
	}
	result = strconv.Itoa(c)
	return result
}

func gap(val string) string {
	const a = '1'
	const b = '0'

	var cur int
	var max int
	max = 0
	cur = 0
	for _, v := range val {
		if v == b {
			cur++
		}
		if cur > max {
			max = cur
		}
		if v == a {
			cur = 0
		}

	}

	return strconv.Itoa(max)

}

func missing(k string) string {
	m := make(map[int]int)
	var r string
	var B []string = strings.Split(k, ",")
	for i, v := range B {
		k, _ := strconv.Atoi(v)
		m[k] = i + 1
	}
	fmt.Println(m)
	for j := 1; j <= len(B)+1; j++ {
		if m[j] == 0 {
			r = strconv.Itoa(j)
		}
	}
	return r
}

func tapeEquilibrium(k string) string {
	var B []string = strings.Split(k, ",")
	m := make(map[int]int)
	var r []int
	var lhs = 0
	var rhs = 0
	for i, v := range B {
		k, _ := strconv.Atoi(v)
		m[i] = k
	}
	for j := 1; j < len(m); j++ {
		for p := 0; p <= len(m); p++ {
			if p < j {
				lhs = lhs + m[p]
			} else {
				rhs = rhs + m[p]
			}
		}
		r = append(r, int(math.Abs(float64(lhs-rhs))))
		fmt.Println(math.Abs(float64(lhs - rhs)))
		lhs = 0
		rhs = 0
	}

	sort.Ints(r)

	return "Equilibrium is " + strconv.Itoa(r[0])
}

func oddOne(k string) string {
	var A []int
	var r int
	var B []string = strings.Split(k, ",")
	for _, s := range B {
		t, _ := strconv.Atoi(s)
		A = append(A, t)
	}
	var l int = len(A)
	for _, p := range A {
		var c int = 0
		fmt.Println(p)

		for _, q := range A {
			if p != q {
				c = c + 1
				fmt.Println("Add")
				fmt.Println(c)
			}

		}
		if c == (l - 1) {
			fmt.Println("Number")
			r = p
		}
	}
	return strconv.Itoa(r)
}

func frogTwo(R string, k int) string {
	//fmt.Println(R, k)
	var A []string = strings.Split(R, ",")
	m := make(map[int]int)
	var n []string
	for i, v := range A {
		j, _ := strconv.Atoi(v)
		m[i] = j - 1
		n = append(n, "sink")
	}
	fmt.Println(m)
	for j := 0; j <= len(m); j++ {
		q := m[j]
		if n[q] != "float" {
			n[q] = "float"
		}
		fmt.Println(n)
		y := 0
		for r := 0; r <= len(m)-1; r++ {
			if n[r] == "float" {
				fmt.Println("Yay! " + strconv.Itoa(r))
				y++
				if y == k {
					return strconv.Itoa(y)
				}
			} else {
				fmt.Println("Sink! ")
				y = 0
			}

		}
	}
	return "Hello"

}

func rotate(R string, k int) string {
	var A []string
	var i int
	A = strings.Split(R, ",")
	B := make([]string, len(A))
	l := len(A)
	if k < 0 {
		k = l - k
	}
	k = k % l
	fmt.Println(k)
	for i = 0; i < l; i++ {
		fmt.Println(i)
		if i+k < l {
			B[i+k] = A[i]

		} else {
			B[k-l+i] = A[i]

		}

	}
	var r string = strings.Join(B, ",")
	return r
}
