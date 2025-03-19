package main
// Mohammad Reyhan Aretha Fatin
// 103112400078
import "fmt"

func faktorBilangan(x int) {
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	faktorBilangan(a)
}
