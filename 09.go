//Fungsi nilai tanpa nilai balik

package main

import (
	"fmt"
)

func tambah(a int, b int) {
	c := a + b
	fmt.Println("Hasil penjumlahan antara", a, "+", b, "Adalah", c)
}
func main() {
	a := 10
	b := 20
	tambah(a, b)
}
