package main

import "fmt"

func main() {

	var nama [3]string

	nama[0] = "Naufal"
	nama[1] = "Yazid"
	nama[2] = "Fakhruddin"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var values = [3]int{
		99,
		69,
		90,
	}
	fmt.Println(values)

}

//array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
//saat membuat array,kita perlu menentukan jumlah data yang bisa di tampung oleh array tersebut
//data tampung array tidak bisa bertambah setelah array tersebut

//function array
//len(array) berfungsii untuk mendapatkan jumlah panjang array
//array[index] befungsi untuk mendapat data di posisi index
//array[index]=value berfungsi untuk mengubah data di posisi index
