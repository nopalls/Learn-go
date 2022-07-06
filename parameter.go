package main

import "fmt"

func sayHelloTo(fistName string, lastName string) {
	fmt.Println("Hello", fistName, lastName)
}

func main() {
	sayHelloTo("Naufal", "Yazid")
}
