// M.DAVI ILYAS RENALDO
// 103112400062
package main

import "fmt"

func fakRekursif(n, i int) {
        if i > n {
                return
        }
        if n%i == 0 {
                fmt.Print(i, " ")
        }
        fakRekursif(n, i+1)
}

func main() {
        var n int
        fmt.Scan(&n)
        fakRekursif(n, 1)
        fmt.Println()
}