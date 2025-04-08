// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func cetakBaris(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	cetakBaris(n-1)
}

func cetakBintang(n int) {
	if n == 0 {
		return
	}
	cetakBintang(n-1)
	cetakBaris(n)
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	
	cetakBintang(n)
}