package main
// Mohammad Reyhan Aretha Fatin
// 103112400078
import "fmt"

func deretFibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return deretFibonacci(x-1)+deretFibonacci(x-2)
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	for i := 0; i <= a; i++ {
		fmt.Print(deretFibonacci(i), " ")
	}
}