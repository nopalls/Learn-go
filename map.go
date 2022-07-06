package main

import "fmt"

func main() {
	person := map[string]string{
		"Name": "Naufal",
		"Asal": "Bekasi",
	}

	fmt.Println(person["Name"])
	fmt.Println(person["Asal"])

	book := make(map[string]string)
	book["Judul"] = "Belajar Golang"
	book["Author"] = "Naufal"
	book["Tebal"] = "Seribu halaman"

	delete(book, "Tebal")
	fmt.Println(book)
}

//function map
//len(map) untuk mendapatkan jumlah data di map
//map[key] mengambil data di map dengan key
//map[key] = value mengubah data di map dengan key
//make(map[TypeKey]TypeValue) untuk membuat map baru
