//RYANAKEYLANOVIANTOWIDODO
//103112400081

package main

import "fmt"

func faktorRecursive(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorRecursive(n, i+i)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	fmt.Printf("Faktor dari %d: ", n)
	faktorRecursive(n, 1)
	fmt.Println()
}
