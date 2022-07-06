package main

import (
	"fmt"
)

func main() {

	var name = "Naufal"

	if name == "Jokowi" {
		fmt.Println("Hello Jokowi")
	} else if name == "Naufal" {
		fmt.Println("Hello Naufal")
	} else {
		fmt.Println("SIpp")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama Kepanjangan")
	} else {
		fmt.Println("OKAYY")
	}
}
