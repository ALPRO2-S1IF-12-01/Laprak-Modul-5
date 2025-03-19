//103112400027_RAJA MUHAMMAD LUFHTI
package main

import"fmt"

func fibonacci(ok int) int {
	if ok <= 1 {
		return ok
	}
	return fibonacci(ok-1) + fibonacci(ok-2)
}

func main() {
	var q int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&q)
	fmt.Print("Fibonacci: ")
	for i := 0; i <= q; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
}