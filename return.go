package main

import "fmt"

func sayHello(name string) string {

	if name == "" {
		return "POV! You're robot"
	} else {
		return "Hello " + name
	}
}

func main() {

	result := sayHello("Naufal")

	fmt.Println(result)

	fmt.Println(sayHello(""))
}

//return multiple values
