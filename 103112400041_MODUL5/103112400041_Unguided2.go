//BERTHA ADELA
//103112400041
package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    hitung(n, 1)
}
func hitung(n, skrg int) int {
    if skrg > n {
        return 0
    }
    cetakBintang(skrg)
    fmt.Println()
    hitung(n, skrg+1)
	return 0
}
func cetakBintang(jumlah int) int {
    if jumlah == 0 {
        return 0
    }
    fmt.Print("*")
    cetakBintang(jumlah-1)
	return 0
}