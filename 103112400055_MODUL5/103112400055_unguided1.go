//Feros Pedrosa

package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai Fibonacci
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int

	// Meminta pengguna untuk memasukkan nilai n
	fmt.Print("Masukkan suku Fibonacci yang diinginkan (n): ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
	}
}
