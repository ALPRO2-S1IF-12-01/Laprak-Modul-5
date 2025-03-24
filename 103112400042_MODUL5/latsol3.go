// M.HANIF AL FAIZ
// 103112400042
package main

import (
	"fmt"
)

func printFactors(n int, current int) {
	if current > n {
		return
	}
	if n%current == 0 {
		fmt.Print(current, " ")
	}
	printFactors(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Faktor dari %d: ", N)
	printFactors(N, 1)
	fmt.Println()
}
