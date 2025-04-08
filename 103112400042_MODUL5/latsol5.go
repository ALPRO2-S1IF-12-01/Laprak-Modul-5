// M.HANIF AL FAIZ
// 103112400042
package main

import (
	"fmt"
)

func cetakBilanganGanjil(n int, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	cetakBilanganGanjil(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Printf("Bilangan ganjil dari 1 hingga %d: ", N)
	cetakBilanganGanjil(N, 1)
	fmt.Println()
}
