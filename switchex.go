package main

import "fmt"

func main() {

	name := "Naufal"

	switch name {
	case "Naufal":
		fmt.Println("Hello Naufal")
	case "Yola":
		fmt.Println("Hello Yola")
	default:
		fmt.Println("Wrong input name")
	}

	//switch dengan kondisi

	switch length := len(name); length > 3 {
	case true:
		fmt.Println("Nama Benar")
	default:
		fmt.Println("Nama Salah")
	}

	//switch tanpa kondisi

	length := len(name)

	switch {
	case length > 3:
		fmt.Println("Nama Benar")
	default:
		fmt.Println("Nama Tidak Benar")
	}

}
