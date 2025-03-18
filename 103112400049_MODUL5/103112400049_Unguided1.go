package main

//HISYAM NURDIATMOKO 103112400049
import "fmt"

func hitungFibonacci(x int) int {
	if x <= 1 {
		return x
	}
	a, b := 0, 1
	for i := 2; i <= x; i++ {
		a, b = b, a+b
	}
	return b
}

func cetakFibonacci(n int) {
	for i := 0; i <= n; i++ {
		fmt.Print(hitungFibonacci(i), " ")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakFibonacci(n)
}
