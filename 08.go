//Membuat fungsi dengan parameter

package main

import (
	"fmt"
)

func tambah(a int, b int) int {
	c := a + b
	return c
}

func main() {
	a := 10
	b := 20

	hasil := tambah(a, b)

	fmt.Println("Hasil Antara", a, "+", b, "Adalah", hasil)
}
