//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func hitungFibonacci(n int) int {
    if n == 0 {
        return 0    
    }
    if n == 1 {
        return 1    
    }
    return hitungFibonacci(n-1) + hitungFibonacci(n-2)
}
func cetakDeretFibonacci(n int) {
    fmt.Print("n  : ")
    for i := 0; i <= n; i++ {
        fmt.Printf("%-4d", i)
    }
    fmt.Println()
    fmt.Print("Sn : ")
    for i := 0; i <= n; i++ {
        fmt.Printf("%-4d", hitungFibonacci(i))
    }
    fmt.Println()
}

func main() {
    cetakDeretFibonacci(10)
}