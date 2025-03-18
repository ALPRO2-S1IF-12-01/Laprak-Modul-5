//Feros Pedrosa

package main

import "fmt"

func faktorRecursive(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorRecursive(n, i+1)
}

func main() {
	var N int

	// Menerima input dari user
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&N)

	fmt.Printf("Faktor dari %d: ", N)
	faktorRecursive(N, 1)
}
