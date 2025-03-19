package main

import "fmt"

// rekursif cetak bintang
func cetakBintang(n int) {
	if n == 0 {
		return
	}

	cetakBintang(n - 1) //buat baris sblmnya

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println() // baris baru
}

func main() {
	var n int

	fmt.Scanln(&n)

	cetakBintang(n)
}
