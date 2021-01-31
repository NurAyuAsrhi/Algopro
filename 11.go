//Array
package main

import (
	"fmt"
)

func main() {
	//contoh penerapan array
	var names [4]string
	names[0] = "nilai indeks 0"
	names[1] = "nilai indeks 1"
	names[2] = "nilai indeks 2"
	names[3] = "nilai indeks 3"

	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println(" ")

	//contoh penerapan array dengan fungsi perulangan
	for i, v := range names {
		fmt.Println("index ke ", i, "=>", v)

	}
}
