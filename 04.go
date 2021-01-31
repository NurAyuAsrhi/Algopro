//contoh break dan continue

package main

import(
	"fmt"
)

func main(){
	for i := 0 i < 10; i++ {
		if i % 2 == 0 {
			continue
		
		fmt.Printf("^%d", i);
	}
	//keluaran dari kode diatas adalah : 13579

	for i := 0; i < 10; i++ {
		if i%2 == 0{
			continue
		}
		fmt.Printf("%d", i);
		if i > 3 {
			break
		}
	}
}
//eluaran dari kode ini 135