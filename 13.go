//Penerapan Pointer

package main

import (
	"fmt"
)

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) //0XC20800A220

	fmt.Println("numberB (value)   : ", *numberB) //4
	fmt.Println("numberB (address) :", &numberA)  //0XC20800A220

}
