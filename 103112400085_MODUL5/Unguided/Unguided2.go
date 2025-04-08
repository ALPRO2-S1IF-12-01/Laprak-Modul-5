// Anastasia Adinda N.I
// 103112400085
package main

import "fmt"

func cetakBaris(n int) {
	if n > 0 {
		fmt.Print("*")
		cetakBaris(n - 1)
	}
}

func cetakBintang(n, i int) {
	if i > n {
		return
	}
	cetakBaris(i)
	fmt.Println()
	cetakBintang(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	cetakBintang(n, 1)
}
