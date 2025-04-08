package main

import (
	"fmt"
)


func cetakBintang(n int, i int) {
	if i > n { 
		return
	}
	fmt.Println(repeat("*", i)) 
	cetakBintang(n, i+1)        
}


func repeat(char string, count int) string {
	if count == 0 {
		return ""
	}
	return char + repeat(char, count-1)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n) 
	cetakBintang(n, 1)
}
