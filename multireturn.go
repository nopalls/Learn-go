package main

import "fmt"

func getFullName() (string, string) {

	return "Naufal", "Yazid"

}

func main() {

	fistName, _ := getFullName()

	fmt.Println(fistName)
	//fmt.Println(lastName)
}

// tanda _ untuk menghiraukan
