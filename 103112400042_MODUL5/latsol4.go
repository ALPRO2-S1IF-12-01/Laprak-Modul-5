// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func printSequence(n int) {

	fmt.Print(n, " ")
	if n > 1 {
		printSequence(n - 1)
	}
	if n != 1 {
		fmt.Print(n, " ")
	}
}

func main() {
	var N int
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&N)
	printSequence(N)
}
