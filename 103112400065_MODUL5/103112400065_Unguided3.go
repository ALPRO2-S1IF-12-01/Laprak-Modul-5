// Dimas Ramadhani
// 103112400065

package main

import "fmt"

func faktorBilangan(x, i int) {
	if i > x {
		return
	}
	if x%i == 0 {
		fmt.Print(i, " ")
	}
	faktorBilangan(x, i+1)
}

func main() {
	var a int
	fmt.Scan(&a)
	faktorBilangan(a, 1)
}
