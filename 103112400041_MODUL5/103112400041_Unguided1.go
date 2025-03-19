//BERTHA ADELA
//103112400041
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	for i:= 1; i<= n; i++ {
		fmt.Print(f(i), " ")
	}
}
func f(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return f(n-1)+f(n-2)
}