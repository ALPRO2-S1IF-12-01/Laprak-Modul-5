//Muhammad Zaky Mubarok
package main

import "fmt"


func cetakBintang(jumlah int, maksimum int) {
	
	if jumlah > maksimum {
		return
	}
	
	for i := 0; i < jumlah; i++ {
		fmt.Print("*")
	}
	
	fmt.Println()
	
	cetakBintang(jumlah+1, maksimum)
}

func main() {
	
	var jumlahBintang int
	fmt.Print("Masukkan jumlah bintang: ")
	fmt.Scanln(&jumlahBintang)

	
	cetakBintang(1, jumlahBintang)
}
