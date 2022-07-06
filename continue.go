package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}

		//jika 1%2 == 0 maka yang keluar adalah hasil ganjil
		//jika 1%2 == 1 maka yang keluar adalah hasil genap
		fmt.Println("Perulangan ke", i)
	}
}
