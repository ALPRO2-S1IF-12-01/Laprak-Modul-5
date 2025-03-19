//103112400027_RAJA MUHAMMAD LUFHTI
package main

import "fmt"

func cetakBintang(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)
	cetakBintang(n)
}