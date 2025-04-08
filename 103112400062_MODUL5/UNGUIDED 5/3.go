// M.DAVI ILYAS RENALDO
// 103112400062
package main

import "fmt"

func pangkatRekursif(x, y int) int {
        if y == 0 {
                return 1
        }
        return x * pangkatRekursif(x, y-1)
}

func main() {
        var x, y int
        fmt.Scan(&x, &y)
        hasil := pangkatRekursif(x, y)
        fmt.Println(hasil)
}