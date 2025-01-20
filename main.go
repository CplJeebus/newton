package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", simple)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}

func simple(w http.ResponseWriter, r *http.Request) {
	const simple = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <title>Binary String Form</title>
</head>
<body>
  <form action="http://localhost:8080/api" method="POST">
    <label for="binaryInput">Enter a string of 0s and 1s:</label>
    <input
      type="text"
      id="gap"
      name="binaryString"
      pattern="[0-1]+"
      required
      placeholder="e.g. 001101"
    />
  </form>
<br>
<form action="http://localhost:8080/api" method="POST">
    <button type="submit">Submit</button>
    <label for="binaryInput">Enter a comma seperated array of stings and an index</label>
    <input
      type="text"
      id="1"
      name="rotate"
      required
      placeholder="e.g. 001101"
    />
	<input
      type="number"
      id="2"
      name="key"
      required
    />
    <button type="submit">Submit</button>

</body>
</html>`

	fmt.Fprint(w, simple)
}

func api(w http.ResponseWriter, r *http.Request) {
	var b string
	_ = r.ParseForm()
	if r.Form.Get("binaryString") != "" {
		b = gap(r.Form.Get("binaryString"))
	} else if r.Form.Get("rotate") != "" {
		var j, _ = strconv.Atoi(r.Form.Get("key"))
		b = rotate(r.Form.Get("rotate"), j)
	}
	fmt.Fprint(w, b)
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

func rotate(R string, k int) string {
	var A []string
	var i int
	A = strings.Split(R, ",")
	B := make([]string, len(A))
	l := len(A)
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
