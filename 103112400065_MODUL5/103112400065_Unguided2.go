// Dimas Ramadhani
// 103112400065

package main

import "fmt"

func buatSegitiga(x int) {
	if x == 0 {
		return
	}
	buatSegitiga(x - 1)
	for i := 1; i <= x; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var a int
	fmt.Scan(&a)
	buatSegitiga(a)
}
