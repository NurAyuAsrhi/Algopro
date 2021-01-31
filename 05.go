//perulngan for-range
package main

import (
	"fmt"
)

func main() {
	data := []string{"panca", "sakti"}

	for index, v := range data {
		fmt.Println("perulangan ke", index, v)
	}
}
