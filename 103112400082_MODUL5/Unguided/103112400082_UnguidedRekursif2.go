package main

import "fmt"

func bintang(j int) {
    if j <= 0 {
        return
    }
    fmt.Print("*")
    bintang(j - 1)
}

func baris(i, tinggi int) {
    if i > tinggi {
        return
    }
    bintang(i)
    fmt.Println()
    baris(i + 1, tinggi)
}

func main() {
    var tinggi int
    fmt.Scan(&tinggi)
    
   baris(1, tinggi)
}