// Raihan Adi Arba
// 103112400071
package main

import "fmt"

func faktorBilangan(x, i int) {
	if i <= x {
		if x%i == 0 {
			fmt.Print(i, " ")
		}
		faktorBilangan(x, i+1)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	faktorBilangan(N, 1)
}
