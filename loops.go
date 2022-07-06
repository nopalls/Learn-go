package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	//for dengan statement

	for counter := 0; counter < 5; counter++ {
		fmt.Println("For ke", counter)
	}
	//counter := 0 adalah init statement
	//counter++ adalah post statement

	//for range
	slice := []string{"Naufal", "Yazid", "Fakhruddin"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for bisa digunakan untuk melakukan interaksi terhadap semua data collection
	//data collection contoh nya array,slice,map dll

	for i, value := range slice {
		fmt.Println("Nama ke", i, "=", value)
	}
}
