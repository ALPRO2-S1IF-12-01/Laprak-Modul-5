//RYANAKEYLANOVIANTOWIDODO
//103112400081

package main

import "fmt"

func fibonacciIteratif(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	var n int
	fmt.Print("Masukkan suku Fibonacci yang diinginkan (n): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i+1, fibonacciIteratif(i))
	}
}
