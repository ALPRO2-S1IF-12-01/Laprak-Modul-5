// Raihan Adi Arba
// 103112400071
package main

import "fmt"

func cetakBintang(n int) {
	if n <= 0 {
		return
	}
	fmt.Print("*")
	cetakBintang(n - 1)
}

func buatSegitiga(x int) {
	if x == 0 {
		return
	}
	buatSegitiga(x - 1)
	cetakBintang(x)
	fmt.Println()
}

func main() {
	var a int
	fmt.Scan(&a)
	buatSegitiga(a)
}
