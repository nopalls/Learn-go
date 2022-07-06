package main

import "fmt"

func main() {
	var nilai32 int32 = 2000
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	var name = "nopal"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
