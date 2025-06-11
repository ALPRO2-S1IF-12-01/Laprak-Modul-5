package main

import (
	"fmt"
	"math"
)

func luasPermukaanTabung(r, t float64) float64 {
	return 2 * math.Pi * r * (r + t)
}

func volumeTabung(r, t float64) float64 {
	return math.Pi * math.Pow(r, 2) * t
}

func main() {
	var r, t float64

	fmt.Print("Masukkan jari-jari tabung: ")
	_, errR := fmt.Scan(&r)
	fmt.Print("Masukkan tinggi tabung: ")
	_, errT := fmt.Scan(&t)
	if errR != nil || errT != nil {
		fmt.Println("Input tidak valid! Harap masukkan angka yang benar.")
		return
	}
	if r <= 0 || t <= 0 {
		fmt.Println("Jari-jari dan tinggi tabung harus lebih dari nol.")
		return
	}
	luas := luasPermukaanTabung(r, t)
	volume := volumeTabung(r, t)
	fmt.Println("===================================")
	fmt.Printf("Luas Permukaan Tabung: %.2f satuan²\n", luas)
	fmt.Printf("Volume Tabung: %.2f satuan³\n", volume)
	fmt.Println("===================================")
}
