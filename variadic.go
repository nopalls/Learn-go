package main

import "fmt"

func tambah(angka ...int) int { //arti ,,, adalah ketika gue manggil func tambah itu gue bisa masukin angka yang banyak

	total := 0

	for _, value := range angka {
		total += value
	}

	return total
}

func main() {

	total := tambah(10, 20, 30, 40, 50) // variable total

	fmt.Println(total)

	//slice parameter
	slice := []int{10, 10, 10, 10, 10}
	total = tambah(slice...)
	fmt.Println(total)
}
