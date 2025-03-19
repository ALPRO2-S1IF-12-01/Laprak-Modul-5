//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func cetakBintang(n int, baris int) {
    if baris > n {
        return
    }
    for i := 1; i <= baris; i++ {
        fmt.Print("*")
    }
    fmt.Println()
    cetakBintang(n, baris+1)
}

func main() {
    var n int

    fmt.Print()
    fmt.Scan(&n)
    cetakBintang(n, 1)
}