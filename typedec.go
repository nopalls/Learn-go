package main

import "fmt"

func main() {

	type NoKTP string
	type Married bool

	var NoKTPNopal NoKTP = "3216080807020012"
	var StatusNikah Married = false

	fmt.Println(NoKTPNopal)
	fmt.Println(StatusNikah)
}
