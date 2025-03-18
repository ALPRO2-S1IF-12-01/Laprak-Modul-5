package main

import (
	"fmt"
)

func printStars(n int) {
	if n == 0 {
		return 
	}
	printStars(n - 1) 
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println() 
}

func main() {
	var n int
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Masukkan angka positif!")
		return
	}

	printStars(n) 
}
