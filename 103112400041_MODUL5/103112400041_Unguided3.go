//BERTHA ADELA
//103112400041
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	prima(n, 1)
}
func prima(p, x int) int {
	if p<x {
		return 0
	}
	if p%x == 0 {
		fmt.Print(x, " ")
	}
	prima(p, x+1)
	return 0
}