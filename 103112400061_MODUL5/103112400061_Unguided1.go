// KEISHIN NAUFA ALFARIDZHI
// 103112400061
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan fibonacci: ")
	fmt.Scan(&n)
	fmt.Printf("Hasil: %v\n", fibonacci(n))
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}