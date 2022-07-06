package main

import "fmt"

func main() {
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[4:8]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = bulan[5:11]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Anjay")
	fmt.Println(slice3)
	slice3[0] = "Not November"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(bulan)

	newSlice := make([]string, 3, 5)

	newSlice[0] = "Naufal"
	newSlice[1] = "yazid"
	newSlice[2] = "Fakruddin"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
