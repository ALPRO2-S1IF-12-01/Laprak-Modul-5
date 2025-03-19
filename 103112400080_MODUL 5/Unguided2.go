// JESIKA METANIA RAHMA ARIFIN
//103112400080package main

package main
import (
	"fmt"
)

func printStars(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	printStars(n - 1)
}
func printPattern(n, i int) {
	if i > n {
		return
	}
	printStars(i)
	fmt.Println()
	printPattern(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	printPattern(n, 1)
}