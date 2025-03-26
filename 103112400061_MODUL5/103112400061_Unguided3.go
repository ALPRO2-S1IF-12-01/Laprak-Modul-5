// KEISHIN NAUFA ALFARIDZHI
// 103112400061
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari bilangan: ")
	faktor(n, 1)
	fmt.Println()
}

func faktor(n int, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		faktor(n, i + 1)
	}
}