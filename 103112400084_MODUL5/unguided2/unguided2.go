//103112400084
//Nufail Alauddin Tsaqif
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	Segitiga(n, 1)
}

func Segitiga(n, i int) {
	if i > n {
		return
	}
	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()
	Segitiga(n, i+1)
}