//Nama: Anggun Wahyu Widiyana (103112480280) 
package main

import "fmt"

func cetakBilangan(n int, angkaSaatIni int) {
    if angkaSaatIni < 1 {
        return
    }
    fmt.Print(angkaSaatIni, " ")
    cetakBilangan(n, angkaSaatIni-1)
    if angkaSaatIni != n {
        fmt.Print(angkaSaatIni, " ")
    }
}

func main() {
    var n int
    fmt.Scan(&n)
    cetakBilangan(n, n)
}