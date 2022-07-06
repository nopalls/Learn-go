package main

import "fmt"

func getNamaLengkap() (firstName, midName, lastName string) {

	firstName = "Naufal"
	midName = "Yazid"
	lastName = "Fakhruddin"

	return
}

func main() {

	firstName, midName, lastName := getNamaLengkap()

	fmt.Println(firstName, midName, lastName)
}
