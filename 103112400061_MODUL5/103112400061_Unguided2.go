// KEISHIN NAUFA ALFARIDZHI
// 103112400061
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	bintang(n)
}

func bintang(n int) {
	if n <= 0 {
		return
	}
	bintang(n-1)
	baris(n)
}

func baris(n int) {
	if n <= 0 {
		fmt.Println()
		return
	}
	fmt.Print("*")
	baris(n-1)
}