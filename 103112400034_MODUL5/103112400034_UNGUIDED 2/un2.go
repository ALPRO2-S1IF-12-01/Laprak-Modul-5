// MULIA AKBAR NANDA PRATAMA
// 103112400034
package main

import "fmt"
func main() {
        var n int
        fmt.Print("masukkan: ")
        fmt.Scan(&n)

        cetakBintang(n)
}
func cetakBintang(n int) {
	if n == 0 {
			return
	}
	cetakBintang(n - 1)
	for i := 0; i < n; i++ {
			fmt.Print("*")
	}
	fmt.Println()
}
