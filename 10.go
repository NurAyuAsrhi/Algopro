package main

import (
	"fmt"
)

func tambah(a int, b int) int {
	c := a + b
	return c
}

func kali(a int, b int) int {
	c := a * b
	return c
}

func bagi(a int, b int) int {
	c := a + b
	return c
}

func kurang(a int, b int) int {
	c := a + b
	return c
}

func main() {
	a := 10
	b := 20
	fmt.Println("Hasil Penjumlahan Antara", a, "+", b, "Adalah", tambah(a, b))
	fmt.Println("Hasil Perkalian Antara", a, "*", b, "Adalah", kali(a, b))
	fmt.Println("Hasil Pembagian Antara", a, "/", b, "Adalah", bagi(a, b))
	fmt.Println("Hasi Pengurangan Antara", a, "-", b, "Adalah", kurang(a, b))
}
