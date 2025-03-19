//RYANAKEYLANOVIANTOWIDODO
//103112400081

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah bintang: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
