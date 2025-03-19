//103112400027_RAJA MUHAMMAD LUFHTI
package main

import (
	"fmt"
)

func cariFaktorial(angka, pembagi int) {
	if pembagi > angka {
		return
	}
	if angka%pembagi == 0 {
		fmt.Print(pembagi, " ")
	}
	cariFaktorial(angka, pembagi+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fmt.Print("Faktor dari ", n, ": ")
	cariFaktorial(n, 1)
	fmt.Println()
}
