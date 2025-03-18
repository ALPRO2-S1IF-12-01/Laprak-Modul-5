package main

//HISYAM NURDIATMOKO 103112400049
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cariFaktor(n, 1)
}

func cariFaktor(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	cariFaktor(n, i+1)
}
