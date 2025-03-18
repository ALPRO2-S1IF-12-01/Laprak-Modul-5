package main
// Mohammad Reyhan Aretha Fatin
// 103112400078
import "fmt"

func cetakSegitiga(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var a int
	fmt.Scan(&a)
	cetakSegitiga(a)
}
