//Nama: Anggun Wahyu Widiyana (103112480280) 
package main

import "fmt"

func cetakGanjil(n int, angka int) {
    if angka > n {
        return
    }
    if angka%2 != 0 {
        fmt.Print(angka, " ")
    }
    cetakGanjil(n, angka+1)
}

func main() {
    var n int
    fmt.Scan(&n)
    cetakGanjil(n, 1)
}