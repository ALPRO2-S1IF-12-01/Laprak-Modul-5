package main
//SETYO NUGGROHO
//103112400024
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