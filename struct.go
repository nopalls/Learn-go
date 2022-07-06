package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var nopal Customer
	nopal.Name = "Naufal Yazid"
	nopal.Address = "Bekasi"
	nopal.Age = 19

	//struct literals

	Joko := Customer{
		Name:    "Joko", //JANGAN LUPA DITAMBAHIN TITIK KOMA DI AKHIR
		Address: "Jakarta",
		Age:     40,
	}

	fmt.Println(nopal)
	fmt.Println(nopal.Name)
	fmt.Println(nopal.Address)

	fmt.Println(Joko)
	fmt.Println(Joko.Name)
	fmt.Println(Joko.Address)

}

//membuat data struct
//struct adalah template data atau prototype data
//struct tidak bisa langsung digunakan
